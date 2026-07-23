package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ideamans/go-llm-cli-kit/llmcmd"
	kitdocs "github.com/ideamans/go-llm-cli-kit/llmdocs"

	"github.com/ideamans/vultr-cli/internal/llmdocs"
)

// TestEmbeddedReference verifies the embedded reference carries the
// hand-written conventions, the generated command catalog and the request body
// schemas — the three things an agent needs to compose a call without guessing.
func TestEmbeddedReference(t *testing.T) {
	g, err := llmdocs.Docs().Markdown()
	if err != nil {
		t.Fatalf("Markdown: %v", err)
	}
	if len(g) == 0 {
		t.Fatal("embedded reference is empty")
	}
	for _, want := range []string{
		"reference for AI agents", // hand-written header
		"VULTR_API_KEY",           // auth convention
		"--all",                   // pagination convention
		"# Command catalog",       // generated catalog
		"## `vultr instance`",     // generated service section
		"vultr instance create <json:InstanceCreateReq>", // generated usage line
		"# Request body schemas",                         // generated schema chapter
		"InstanceCreateReq: `{",                          // generated schema entry
		"DomainRecordCreateReq",                          // another known body type
	} {
		if !strings.Contains(g, want) {
			t.Errorf("embedded reference missing %q", want)
		}
	}
}

// TestChapterOrder pins the reading order: an agent must meet the conventions
// before the catalog, or it will compose calls from signatures alone.
func TestChapterOrder(t *testing.T) {
	sections, err := llmdocs.Docs().Sections()
	if err != nil {
		t.Fatalf("Sections: %v", err)
	}
	var files []string
	for _, s := range sections {
		files = append(files, s.File)
	}
	want := []string{"00-guide.md", "90-commands.md", "91-schemas.md"}
	if strings.Join(files, ",") != strings.Join(want, ",") {
		t.Errorf("chapters = %v, want %v", files, want)
	}
}

// TestGeneratedCatalogIsCurrent fails when the committed catalog no longer
// matches the command tree — the same check CI runs via git diff, but with a
// message that names the fix.
func TestGeneratedCatalogIsCurrent(t *testing.T) {
	sections, err := llmdocs.Docs().Sections()
	if err != nil {
		t.Fatalf("Sections: %v", err)
	}
	var committed string
	for _, s := range sections {
		if s.File == "91-schemas.md" {
			committed = s.Body
		}
	}
	if strings.TrimSpace(committed) != strings.TrimSpace(bodySchemasMarkdown()) {
		t.Error("internal/llmdocs/91-schemas.md is stale — run `go generate ./...` and commit the result")
	}
}

func TestLLMSubcommand(t *testing.T) {
	root := newRoot()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"llm"})
	if err := root.Execute(); err != nil {
		t.Fatalf("vultr llm: %v", err)
	}
	if !strings.Contains(out.String(), "# Command catalog") {
		t.Errorf("vultr llm did not print the reference:\n%s", out.String()[:min(400, out.Len())])
	}
}

func TestLLMSubcommandJSON(t *testing.T) {
	root := newRoot()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"llm", "--format", "json"})
	if err := root.Execute(); err != nil {
		t.Fatalf("vultr llm --format json: %v", err)
	}
	var sections []kitdocs.Section
	if err := json.Unmarshal(out.Bytes(), &sections); err != nil {
		t.Fatalf("output is not JSON: %v", err)
	}
	if len(sections) != 3 {
		t.Errorf("got %d sections, want 3", len(sections))
	}
}

// TestLegacyLLMFlag guards the compatibility promise: --llm used to work at any
// position on the command line, and callers still rely on it.
func TestLegacyLLMFlag(t *testing.T) {
	for _, args := range [][]string{
		{"--llm"},
		{"instance", "list", "--llm"},
	} {
		var out bytes.Buffer
		handled, err := llmcmd.HandleLegacy(args, llmConfig(), &out)
		if err != nil {
			t.Fatalf("HandleLegacy(%v): %v", args, err)
		}
		if !handled {
			t.Errorf("HandleLegacy(%v) did not handle --llm", args)
		}
		if !strings.Contains(out.String(), "VULTR_API_KEY") {
			t.Errorf("HandleLegacy(%v) printed the wrong thing", args)
		}
	}
}
