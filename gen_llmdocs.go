package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ideamans/go-llm-cli-kit/catalog"
	"github.com/spf13/cobra"
)

// docsDir is where the embedded reference chapters live.
const docsDir = "internal/llmdocs"

// newGenerateCommand returns the hidden command that regenerates the derived
// chapters of the embedded reference.
//
// It lives in package main rather than in its own internal/gen-llmdocs program
// because generating the catalog needs the assembled cobra tree and the govultr
// reflection helpers, both of which are defined here. Extracting them would be a
// wholesale restructuring of a working CLI for no benefit at the point of use.
func newGenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use:    "gen-llmdocs",
		Short:  "regenerate the embedded LLM reference (development only)",
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			chapters := map[string]string{
				"90-commands.md": catalog.Markdown(newRoot(), catalog.Options{
					Title: "Command catalog",
					Intro: "Generated from the govultr service interfaces by `go generate ./...`.\n" +
						"Do not edit by hand. Every operation is listed as\n" +
						"`vultr <service> <operation> <args>` followed by its govultr signature.",
					// llm documents itself in 00-guide.md; gen-llmdocs is a
					// development command and must not be advertised to agents.
					Compact: true,
					Skip:    []string{"llm", "gen-llmdocs"},
				}),
				"91-schemas.md": bodySchemasMarkdown(),
			}

			for name, content := range chapters {
				path := filepath.Join(docsDir, name)
				if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
					return err
				}
				fmt.Fprintf(cmd.ErrOrStderr(), "wrote %s (%d bytes)\n", path, len(content))
			}
			return nil
		},
	}
}
