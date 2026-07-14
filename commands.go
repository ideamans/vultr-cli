package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vultr/govultr/v3"
)

var (
	contextType      = reflect.TypeOf((*context.Context)(nil)).Elem()
	errorType        = reflect.TypeOf((*error)(nil)).Elem()
	httpResponseType = reflect.TypeOf((*http.Response)(nil))
	listOptionsType  = reflect.TypeOf((*govultr.ListOptions)(nil))
	metaType         = reflect.TypeOf((*govultr.Meta)(nil))
)

// buildCommands reflects over every service interface on govultr.Client and
// registers one subcommand per service, with one sub-subcommand per API method.
func buildCommands(root *cobra.Command) {
	ct := reflect.TypeOf(govultr.Client{})
	for i := 0; i < ct.NumField(); i++ {
		f := ct.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.Interface {
			continue
		}
		svcCmd := &cobra.Command{
			Use:   kebab(f.Name),
			Short: fmt.Sprintf("%s API (%d operations)", f.Name, f.Type.NumMethod()),
		}
		for j := 0; j < f.Type.NumMethod(); j++ {
			svcCmd.AddCommand(newMethodCommand(f.Name, f.Type.Method(j)))
		}
		root.AddCommand(svcCmd)
	}
}

func newMethodCommand(fieldName string, m reflect.Method) *cobra.Command {
	mt := m.Type

	use := kebab(m.Name)
	hasListOpts := false
	hasJSONParam := false
	for k := 0; k < mt.NumIn(); k++ {
		pt := mt.In(k)
		if pt == contextType {
			continue
		}
		if pt == listOptionsType {
			hasListOpts = true
			continue
		}
		use += " " + placeholder(pt)
		if isJSONParam(pt) {
			hasJSONParam = true
		}
	}

	cmd := &cobra.Command{
		Use:                   use,
		Short:                 signatureString(m),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMethod(cmd, fieldName, m, args)
		},
	}
	if hasListOpts {
		fl := cmd.Flags()
		fl.Int("per-page", 0, "number of items per page")
		fl.String("cursor", "", "pagination cursor")
		fl.String("main-ip", "", "filter by main IP (where supported)")
		fl.String("label", "", "filter by label (where supported)")
		fl.String("tag", "", "filter by tag (where supported)")
		fl.String("region", "", "filter by region (where supported)")
		fl.String("description", "", "filter by description (where supported)")
		if returnsMeta(mt) {
			fl.Bool("all", false, "automatically follow pagination and return all items")
		}
	}
	if hasJSONParam {
		cmd.Flags().Bool("schema", false, "print the expected JSON body structure and exit")
	}
	return cmd
}

func isJSONParam(pt reflect.Type) bool {
	switch pt.Kind() {
	case reflect.Ptr:
		return pt.Elem().Kind() == reflect.Struct
	case reflect.Struct, reflect.Map:
		return true
	case reflect.Slice:
		return pt.Elem().Kind() != reflect.String
	}
	return false
}

func placeholder(pt reflect.Type) string {
	switch pt.Kind() {
	case reflect.String:
		return "<string>"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "<int>"
	case reflect.Bool:
		return "<bool>"
	case reflect.Ptr:
		switch pt.Elem().Kind() {
		case reflect.Bool:
			return "[<bool|null>]"
		case reflect.String:
			return "[<string|null>]"
		case reflect.Int, reflect.Int64:
			return "[<int|null>]"
		}
		return "<json:" + typeName(pt.Elem().String()) + ">"
	case reflect.Slice:
		if pt.Elem().Kind() == reflect.String {
			return "<v1,v2,...>"
		}
		return "<json:" + typeName(pt.String()) + ">"
	default:
		return "<json:" + typeName(pt.String()) + ">"
	}
}

func returnsMeta(mt reflect.Type) bool {
	for k := 0; k < mt.NumOut(); k++ {
		if mt.Out(k) == metaType {
			return true
		}
	}
	return false
}

func signatureString(m reflect.Method) string {
	mt := m.Type
	var params, rets []string
	for k := 0; k < mt.NumIn(); k++ {
		if mt.In(k) == contextType {
			continue
		}
		params = append(params, typeName(mt.In(k).String()))
	}
	for k := 0; k < mt.NumOut(); k++ {
		ot := mt.Out(k)
		if ot == httpResponseType || ot == errorType {
			continue
		}
		rets = append(rets, typeName(ot.String()))
	}
	s := fmt.Sprintf("%s(%s)", m.Name, strings.Join(params, ", "))
	if len(rets) > 0 {
		s += " -> " + strings.Join(rets, ", ")
	}
	return s
}

func runMethod(cmd *cobra.Command, fieldName string, m reflect.Method, args []string) error {
	mt := m.Type

	if f := cmd.Flags().Lookup("schema"); f != nil && f.Changed {
		return printSchemas(mt)
	}

	client, err := newClient()
	if err != nil {
		return err
	}
	svc := reflect.ValueOf(client).Elem().FieldByName(fieldName)
	fn := svc.MethodByName(m.Name)

	in := []reflect.Value{reflect.ValueOf(context.Background())}
	ai := 0
	var listOpts *govultr.ListOptions
	for k := 0; k < mt.NumIn(); k++ {
		pt := mt.In(k)
		switch {
		case pt == contextType:
			continue
		case pt == listOptionsType:
			listOpts = buildListOptions(cmd)
			in = append(in, reflect.ValueOf(listOpts))
		default:
			v, err := consumeArg(pt, args, &ai)
			if err != nil {
				return err
			}
			in = append(in, v)
		}
	}
	if ai < len(args) {
		return fmt.Errorf("too many arguments: unexpected %q", args[ai])
	}

	printable, err := call(fn, in)
	if err != nil {
		return err
	}

	all := false
	if f := cmd.Flags().Lookup("all"); f != nil {
		all, _ = cmd.Flags().GetBool("all")
	}
	if all && len(printable) == 2 && printable[0].Kind() == reflect.Slice && printable[1].Type() == metaType {
		combined := printable[0]
		meta, _ := printable[1].Interface().(*govultr.Meta)
		for meta != nil && meta.Links != nil && meta.Links.Next != "" {
			listOpts.Cursor = meta.Links.Next
			page, err := call(fn, in)
			if err != nil {
				return err
			}
			combined = reflect.AppendSlice(combined, page[0])
			meta, _ = page[1].Interface().(*govultr.Meta)
		}
		return printJSON(map[string]any{
			"data": combined.Interface(),
			"meta": map[string]any{"total": combined.Len()},
		})
	}

	return printResult(printable)
}

// call invokes the API method and returns the printable results
// (everything except the *http.Response and the trailing error).
func call(fn reflect.Value, in []reflect.Value) ([]reflect.Value, error) {
	outs := fn.Call(in)
	last := outs[len(outs)-1]
	if last.Type() == errorType && !last.IsNil() {
		return nil, last.Interface().(error)
	}
	var printable []reflect.Value
	for _, o := range outs[:len(outs)-1] {
		if o.Type() == httpResponseType {
			continue
		}
		printable = append(printable, o)
	}
	return printable, nil
}

// printResult renders API return values. A single value prints as-is; with
// multiple values the first becomes "data" and the rest are keyed by their
// type name ("meta" for *govultr.Meta), e.g. Database.ListConnectionPools
// -> {"data": ..., "database_connection_pools": [...], "meta": {...}}.
func printResult(printable []reflect.Value) error {
	switch len(printable) {
	case 0:
		return nil
	case 1:
		return printJSON(printable[0].Interface())
	}
	out := map[string]any{"data": printable[0].Interface()}
	for _, p := range printable[1:] {
		key := "meta"
		if p.Type() != metaType {
			key = resultKey(p.Type())
		}
		out[key] = p.Interface()
	}
	return printJSON(out)
}

// resultKey derives a snake_case JSON key from a return type, pluralizing
// slices: []govultr.AvailableOption -> "available_options".
func resultKey(t reflect.Type) string {
	plural := false
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		if t.Kind() == reflect.Slice {
			plural = true
		}
		t = t.Elem()
	}
	key := strings.ReplaceAll(kebab(t.Name()), "-", "_")
	if plural && !strings.HasSuffix(key, "s") {
		key += "s"
	}
	return key
}

func buildListOptions(cmd *cobra.Command) *govultr.ListOptions {
	fl := cmd.Flags()
	opts := &govultr.ListOptions{}
	opts.PerPage, _ = fl.GetInt("per-page")
	opts.Cursor, _ = fl.GetString("cursor")
	opts.MainIP, _ = fl.GetString("main-ip")
	opts.Label, _ = fl.GetString("label")
	opts.Tag, _ = fl.GetString("tag")
	opts.Region, _ = fl.GetString("region")
	opts.Description, _ = fl.GetString("description")
	return opts
}

func printSchemas(mt reflect.Type) error {
	schemas := map[string]any{}
	for k := 0; k < mt.NumIn(); k++ {
		pt := mt.In(k)
		if pt == contextType || pt == listOptionsType || !isJSONParam(pt) {
			continue
		}
		base := pt
		if base.Kind() == reflect.Ptr {
			base = base.Elem()
		}
		schemas[typeName(base.String())] = schemaOf(pt, 0)
	}
	if len(schemas) == 1 {
		for _, s := range schemas {
			return printJSON(s)
		}
	}
	return printJSON(schemas)
}
