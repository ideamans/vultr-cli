// Package llmdocs embeds the reference that `vultr llm` prints.
//
// 00-guide.md is hand-written. 90-commands.md and 91-schemas.md are generated
// from the govultr reflection data by `go generate ./...` (see gen_llmdocs.go)
// and committed, because go:embed needs real files at build time. CI
// regenerates them and fails if the committed copies are stale.
package llmdocs

import (
	"embed"

	kit "github.com/ideamans/go-llm-cli-kit/llmdocs"
)

//go:embed *.md
var files embed.FS

// Docs is the embedded reference bundle.
func Docs() *kit.Docs { return kit.New(files, ".") }
