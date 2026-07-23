# Generated artifacts — do not hand-edit

These files are derived. Editing them directly is always wrong: the next
`go generate ./...` overwrites the edit, and CI fails on the stale diff in the
meantime.

| Generated file | Source of truth |
| --- | --- |
| `internal/llmdocs/90-commands.md` | the cobra tree built by `buildCommands` in `commands.go`, rendered by `newGenerateCommand` in `gen_llmdocs.go` |
| `internal/llmdocs/91-schemas.md` | the govultr request types, via `bodySchemasMarkdown` in `llm.go` |

Hand-written and safe to edit:

- `internal/llmdocs/00-guide.md` — conventions, auth, workflows, failure modes
- `plugins/vultr-cli/skills/*/SKILL.md` — distributed Agent Skills
- `context7.json` — the pitfall rules

To regenerate: `/regen-ai`, or `go generate ./... && go test ./...`.
