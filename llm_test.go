package main

import (
	"strings"
	"testing"
)

// TestLLMHelp verifies the LLM guide contains the hand-written conventions,
// the full generated catalog, and the request body schemas.
func TestLLMHelp(t *testing.T) {
	g := llmHelp()
	if len(g) == 0 {
		t.Fatal("llmHelp() is empty")
	}
	for _, want := range []string{
		"Guide for LLMs", // hand-written header
		"VULTR_API_KEY",  // auth convention
		"### instance (", // generated service section
		"vultr instance create <json:InstanceCreateReq>", // generated usage line
		"--all",                   // pagination convention
		"## Request body schemas", // schema section
		"InstanceCreateReq: `{",   // generated schema entry
		"DomainRecordCreateReq",   // another known body type
	} {
		if !strings.Contains(g, want) {
			t.Errorf("llmHelp() missing %q", want)
		}
	}
}
