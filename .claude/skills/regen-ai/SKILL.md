---
name: regen-ai
description: Regenerate the embedded LLM reference and verify the result. Use after changing commands, flags, the hand-written guide, or the govultr dependency.
allowed-tools: Bash(go generate:*) Bash(go test:*) Bash(go build:*) Bash(git status:*) Bash(git diff:*) Read
---

# regen-ai

Bring `internal/llmdocs/` back in line with the code.

1. `git status --short` — note what is already dirty, so the regeneration diff
   can be told apart from the user's work in progress.
2. `go generate ./...` — rewrites `90-commands.md` and `91-schemas.md`.
3. `go build ./... && go test ./...`.
4. Report: which chapters changed, how many lines, and whether anything in the
   command catalog changed shape (a new service, a removed operation, a changed
   signature). A large unexpected diff usually means the govultr dependency
   moved — say so rather than committing it silently.

This skill is Claude Code-local; it is not part of the distributed plugin.
