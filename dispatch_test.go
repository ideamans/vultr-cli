package main

import (
	"reflect"
	"testing"

	"github.com/vultr/govultr/v3"
)

// TestNaming pins the kebab-case conversion for tricky govultr identifiers.
func TestNaming(t *testing.T) {
	cases := map[string]string{
		"Account":                  "account",
		"BareMetalServer":          "bare-metal-server",
		"SSHKey":                   "ssh-key",
		"ISO":                      "iso",
		"OS":                       "os",
		"OIDC":                     "oidc",
		"VPC2":                     "vpc2",
		"VirtualFileSystemStorage": "virtual-file-system-storage",
		"CreateIPv4":               "create-ipv4",
		"ListIPv6":                 "list-ipv6",
		"DefaultReverseIPv4":       "default-reverse-ipv4",
		"ISOStatus":                "iso-status",
		"ListVPC2Info":             "list-vpc2-info",
		"ListVPCInfo":              "list-vpc-info",
		"MassStart":                "mass-start",
		"GetKubeConfig":            "get-kube-config",
	}
	for in, want := range cases {
		if got := kebab(in); got != want {
			t.Errorf("kebab(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestDispatcherAssumptions verifies, for every method of every govultr
// service, the signature invariants the reflection dispatcher relies on.
// If a govultr upgrade breaks any of these, this test fails instead of the
// CLI misbehaving at runtime.
func TestDispatcherAssumptions(t *testing.T) {
	ct := reflect.TypeOf(govultr.Client{})
	services, methods := 0, 0

	for i := 0; i < ct.NumField(); i++ {
		f := ct.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.Interface {
			continue
		}
		services++
		seen := map[string]string{}
		for j := 0; j < f.Type.NumMethod(); j++ {
			m := f.Type.Method(j)
			mt := m.Type
			methods++
			id := f.Name + "." + m.Name

			if mt.NumIn() == 0 || mt.In(0) != contextType {
				t.Errorf("%s: first param is not context.Context: %s", id, mt)
			}
			if mt.IsVariadic() {
				t.Errorf("%s: variadic methods are not supported by the dispatcher", id)
			}
			if mt.NumOut() == 0 || mt.Out(mt.NumOut()-1) != errorType {
				t.Errorf("%s: last return value is not error: %s", id, mt)
			}

			listOpts := 0
			for k := 1; k < mt.NumIn(); k++ {
				pt := mt.In(k)
				if pt == contextType {
					t.Errorf("%s: unexpected context.Context at param %d", id, k)
				}
				if pt == listOptionsType {
					listOpts++
					continue
				}
				if !supportedParam(pt) {
					t.Errorf("%s: param %d has unsupported type %s", id, k, pt)
				}
			}
			if listOpts > 1 {
				t.Errorf("%s: %d *ListOptions params", id, listOpts)
			}

			for k := 0; k < mt.NumOut()-1; k++ {
				if mt.Out(k) == errorType {
					t.Errorf("%s: error at non-final return position %d", id, k)
				}
			}

			// catalog printable results; if Meta is returned it must come
			// last so printResult labels it "meta"
			printable := 0
			metaLast := false
			for k := 0; k < mt.NumOut()-1; k++ {
				if mt.Out(k) == httpResponseType {
					continue
				}
				metaLast = mt.Out(k) == metaType
				printable++
			}
			if printable > 3 {
				t.Errorf("%s: too many return values for printResult: %s", id, mt)
			}
			hasMeta := false
			for k := 0; k < mt.NumOut()-1; k++ {
				if mt.Out(k) == metaType {
					hasMeta = true
				}
			}
			if hasMeta && !metaLast {
				t.Errorf("%s: *Meta is not the last printable return: %s", id, mt)
			}

			// the --all pagination contract: a method taking *ListOptions and
			// returning *Meta must have exactly the (data, meta) shape
			if listOpts == 1 && hasMeta && printable != 2 {
				t.Errorf("%s: --all expects (data, *Meta) but got %s", id, mt)
			}

			kb := kebab(m.Name)
			if prev, dup := seen[kb]; dup {
				t.Errorf("%s: command name %q collides with method %s", id, kb, prev)
			}
			seen[kb] = m.Name
		}
	}

	t.Logf("verified %d services, %d methods", services, methods)
	if services < 30 || methods < 400 {
		t.Errorf("suspiciously low API surface: services=%d methods=%d", services, methods)
	}
}

// supportedParam mirrors the type coverage of parseValue in args.go.
func supportedParam(pt reflect.Type) bool {
	switch pt.Kind() {
	case reflect.String, reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Slice, reflect.Map, reflect.Struct:
		return true
	case reflect.Ptr:
		switch pt.Elem().Kind() {
		case reflect.Struct, reflect.String, reflect.Bool,
			reflect.Int, reflect.Int64, reflect.Float64:
			return true
		}
	}
	return false
}
