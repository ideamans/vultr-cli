package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// consumeArg converts the next positional argument to the given parameter type.
// Trailing pointer/slice/map parameters may be omitted and become nil.
func consumeArg(pt reflect.Type, args []string, ai *int) (reflect.Value, error) {
	if *ai >= len(args) {
		switch pt.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Map:
			return reflect.Zero(pt), nil
		}
		return reflect.Value{}, fmt.Errorf("missing required argument of type %s", typeName(pt.String()))
	}
	raw := args[*ai]
	*ai++
	return parseValue(pt, raw)
}

func parseValue(pt reflect.Type, raw string) (reflect.Value, error) {
	switch pt.Kind() {
	case reflect.String:
		return reflect.ValueOf(raw).Convert(pt), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("expected an integer, got %q", raw)
		}
		v := reflect.New(pt).Elem()
		v.SetInt(n)
		return v, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("expected an unsigned integer, got %q", raw)
		}
		v := reflect.New(pt).Elem()
		v.SetUint(n)
		return v, nil
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("expected a number, got %q", raw)
		}
		v := reflect.New(pt).Elem()
		v.SetFloat(f)
		return v, nil
	case reflect.Bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("expected true/false, got %q", raw)
		}
		return reflect.ValueOf(b), nil
	case reflect.Ptr:
		if raw == "null" {
			return reflect.Zero(pt), nil
		}
		switch pt.Elem().Kind() {
		case reflect.String, reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			ev, err := parseValue(pt.Elem(), raw)
			if err != nil {
				return reflect.Value{}, err
			}
			p := reflect.New(pt.Elem())
			p.Elem().Set(ev)
			return p, nil
		}
		return parseJSONValue(pt, raw)
	case reflect.Slice:
		// []string accepts comma-separated values or a JSON array
		if pt.Elem().Kind() == reflect.String && !strings.HasPrefix(strings.TrimSpace(raw), "[") {
			parts := strings.Split(raw, ",")
			v := reflect.MakeSlice(pt, len(parts), len(parts))
			for i, p := range parts {
				v.Index(i).SetString(strings.TrimSpace(p))
			}
			return v, nil
		}
		return parseJSONValue(pt, raw)
	default:
		return parseJSONValue(pt, raw)
	}
}

// parseJSONValue decodes a JSON argument into the parameter type.
// The argument may be inline JSON, "@path/to/file.json", or "-" for stdin.
func parseJSONValue(pt reflect.Type, raw string) (reflect.Value, error) {
	data, err := resolveData(raw)
	if err != nil {
		return reflect.Value{}, err
	}
	target := pt
	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}
	v := reflect.New(target)
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v.Interface()); err != nil {
		return reflect.Value{}, fmt.Errorf("invalid JSON for %s: %w (use --schema to see expected fields)", typeName(pt.String()), err)
	}
	if pt.Kind() == reflect.Ptr {
		return v, nil
	}
	return v.Elem(), nil
}

func resolveData(raw string) ([]byte, error) {
	switch {
	case raw == "-":
		return io.ReadAll(os.Stdin)
	case strings.HasPrefix(raw, "@"):
		return os.ReadFile(raw[1:])
	default:
		return []byte(raw), nil
	}
}

// schemaOf builds a JSON skeleton describing a request struct, using field
// json tags with the type name as placeholder value.
func schemaOf(t reflect.Type, depth int) any {
	if depth > 8 {
		return "..."
	}
	switch t.Kind() {
	case reflect.Ptr:
		return schemaOf(t.Elem(), depth)
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "int"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Slice, reflect.Array:
		return []any{schemaOf(t.Elem(), depth+1)}
	case reflect.Map:
		return map[string]any{"<key>": schemaOf(t.Elem(), depth+1)}
	case reflect.Interface:
		return "any"
	case reflect.Struct:
		if t.PkgPath() == "time" && t.Name() == "Time" {
			return "string (RFC3339)"
		}
		out := map[string]any{}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if !f.IsExported() {
				continue
			}
			if f.Anonymous {
				if inner, ok := schemaOf(f.Type, depth).(map[string]any); ok {
					for k, v := range inner {
						out[k] = v
					}
					continue
				}
			}
			name := strings.Split(f.Tag.Get("json"), ",")[0]
			if name == "-" {
				continue
			}
			if name == "" {
				name = f.Name
			}
			out[name] = schemaOf(f.Type, depth+1)
		}
		return out
	default:
		return t.Kind().String()
	}
}
