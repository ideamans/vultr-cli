# Publishing the vultr-cli plugin

## Before every release

1. `go generate ./...` — regenerate `internal/llmdocs/90-commands.md` and
   `91-schemas.md`, and commit any diff.
2. `go test ./...` — includes `TestPluginSkills`, which enforces that
   `plugin.json.version` equals `pluginVersion` in `main.go` and that the
   SKILL.md frontmatter stays within the Agent Skills standard.
3. `claude plugin validate plugins/vultr-cli` — Claude Code's own validator.
4. Bump `pluginVersion` in `main.go` and `version` in
   `.claude-plugin/plugin.json` together, in the same commit as the release tag.
   The release workflow refuses a tag that disagrees with `pluginVersion`.

## Registering in the marketplace (first release only)

Add to `.claude-plugin/marketplace.json` in `ideamans/claude-public-plugins`:

```json
{
  "name": "vultr-cli",
  "source": {
    "source": "git-subdir",
    "url": "https://github.com/ideamans/vultr-cli.git",
    "path": "plugins/vultr-cli"
  }
}
```

`git-subdir` points at this repository's default branch, so later skill updates
reach users on merge — the marketplace entry only needs touching when the
description changes.

## Verifying the published result

```
/plugin marketplace add ideamans/claude-public-plugins
/plugin install vultr-cli@ideamans-plugins
/vultr-cli-usage
```

Other hosts install the same files directly:

```bash
gh skill install ideamans/vultr-cli/plugins/vultr-cli/skills/vultr-cli-usage --agent copilot
gh skill update
```
