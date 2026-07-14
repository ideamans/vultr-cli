package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/vultr/govultr/v3"
)

// llmHelp renders a single self-contained reference for AI (LLM) consumption:
// usage conventions, the full service/operation catalog, and the JSON schema
// of every request body type. Generated from the same reflection data as the
// command tree, so it never drifts from the actual CLI.
func llmHelp() string {
	var b strings.Builder

	b.WriteString(`# vultr CLI — Guide for LLMs

An aws-style CLI wrapping the entire Vultr API (via govultr). Invocation:

    vultr <service> <operation> [args...] [flags]

## Conventions

- Authentication: set the VULTR_API_KEY environment variable, or pass --api-key <key>.
- If the API key is IP-restricted and requests fail intermittently with 401, add -4 (--ipv4) to force IPv4.
- Arguments are positional and typed by the operation signature shown in the catalog below:
  - <string> <int> <bool>: plain positional values.
  - <json:TypeName>: a JSON request body. Pass inline JSON, @path/to/file.json, or - to read stdin.
    The full schema of every TypeName is listed in the "Request body schemas" section.
    Unknown JSON fields are rejected. Running the operation with --schema prints the schema and exits.
  - <v1,v2,...>: a []string; pass comma-separated values or a JSON array.
  - [<bool|null>] and similar: optional trailing argument; omit it or pass the literal null.
- Operations whose signature includes *ListOptions accept pagination/filtering flags instead of a
  positional argument: --per-page <n>, --cursor <c>, and where supported --tag, --label, --region,
  --main-ip, --description. List operations returning Meta also accept --all to auto-follow
  pagination and merge every page.
- Output is JSON on stdout. Single-value results print as-is. Multi-value results print as an
  object: {"data": ..., "meta": ...} (extra values are keyed by their snake_case type name).
  Operations without a return value print nothing and exit 0. Errors go to stderr, exit code 1.
- Some list operations take a required string filter first; pass an empty string '' for "all"
  (e.g. vultr plan list '' --per-page 5).

## Services and operations

`)

	bodyTypes := map[string]reflect.Type{}
	ct := reflect.TypeOf(govultr.Client{})
	for i := 0; i < ct.NumField(); i++ {
		f := ct.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.Interface {
			continue
		}
		fmt.Fprintf(&b, "### %s (%d operations)\n\n", kebab(f.Name), f.Type.NumMethod())
		for j := 0; j < f.Type.NumMethod(); j++ {
			m := f.Type.Method(j)
			mt := m.Type
			usage := "vultr " + kebab(f.Name) + " " + kebab(m.Name)
			flags := ""
			for k := 1; k < mt.NumIn(); k++ {
				pt := mt.In(k)
				if pt == listOptionsType {
					flags = " [--per-page N --cursor C ...]"
					if returnsMeta(mt) {
						flags = " [--per-page N --cursor C --all ...]"
					}
					continue
				}
				usage += " " + placeholder(pt)
				collectBodyType(pt, bodyTypes)
			}
			fmt.Fprintf(&b, "- `%s%s` — %s\n", usage, flags, signatureString(m))
		}
		b.WriteString("\n")
	}

	b.WriteString("## Request body schemas\n\n")
	b.WriteString("Field values indicate the expected JSON type. All fields are optional unless the API rejects their absence.\n\n")
	names := make([]string, 0, len(bodyTypes))
	for name := range bodyTypes {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		schema, err := json.Marshal(schemaOf(bodyTypes[name], 0))
		if err != nil {
			continue
		}
		fmt.Fprintf(&b, "- %s: `%s`\n", name, schema)
	}
	b.WriteString("\n")
	return b.String()
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
