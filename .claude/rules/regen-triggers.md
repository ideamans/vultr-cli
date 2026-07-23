---
paths:
  - "commands.go"
  - "args.go"
  - "naming.go"
  - "main.go"
  - "llm.go"
  - "gen_llmdocs.go"
  - "internal/llmdocs/00-guide.md"
---

# You just touched the source of the embedded LLM reference

Run `/regen-ai` before finishing, so `internal/llmdocs/90-commands.md` and
`91-schemas.md` match the command tree. CI regenerates them and fails on a dirty
tree, so skipping this only moves the failure later.

Do not edit the generated chapters directly.
