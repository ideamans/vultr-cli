package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/vultr/govultr/v3"
)

// bodySchemasMarkdown renders the JSON schema of every struct type that can be
// passed as a request body, as one chapter of the embedded LLM reference.
//
// The command catalog names these types in <json:TypeName> placeholders but
// cannot describe their shape; without this chapter an agent has no way to
// compose a body except by guessing, and unknown fields are rejected.
func bodySchemasMarkdown() string {
	var b strings.Builder

	b.WriteString(`# Request body schemas

Generated from the govultr request types by ` + "`go generate ./...`" + `. Do not edit by hand.

Every ` + "`<json:TypeName>`" + ` placeholder in the command catalog resolves to one of the
types below. Field values indicate the expected JSON type. All fields are optional
unless the API rejects their absence, and unknown fields are rejected outright —
run the operation with ` + "`--schema`" + ` to print the schema for that operation alone.

`)

	bodyTypes := map[string]reflect.Type{}
	ct := reflect.TypeOf(govultr.Client{})
	for i := 0; i < ct.NumField(); i++ {
		f := ct.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.Interface {
			continue
		}
		for j := 0; j < f.Type.NumMethod(); j++ {
			mt := f.Type.Method(j).Type
			for k := 1; k < mt.NumIn(); k++ {
				if mt.In(k) == listOptionsType {
					continue
				}
				collectBodyType(mt.In(k), bodyTypes)
			}
		}
	}

	names := make([]string, 0, len(bodyTypes))
	for name := range bodyTypes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		schema, err := marshalSchema(schemaOf(bodyTypes[name], 0))
		if err != nil {
			continue
		}
		fmt.Fprintf(&b, "- %s: `%s`\n", name, schema)
	}

	return b.String()
}

// marshalSchema encodes a schema without HTML escaping, so map keys read as
// `{"<key>": ...}` rather than `{"\u003ckey\u003e": ...}`.
func marshalSchema(v any) (string, error) {
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return "", err
	}
	return strings.TrimRight(buf.String(), "\n"), nil
}

// collectBodyType records struct types that are passed as JSON positional
// arguments, keyed by their bare type name as shown in <json:...> placeholders.
func collectBodyType(pt reflect.Type, into map[string]reflect.Type) {
	t := pt
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct || t.PkgPath() == "time" {
		return
	}
	into[typeName(t.String())] = t
}
