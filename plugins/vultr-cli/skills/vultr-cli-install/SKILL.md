---
name: vultr-cli-install
description: Make the vultr command available, installing it only if it is missing. Use when another skill reports that `vultr` is not on PATH, or when the user asks to install, update or upgrade the ideamans vultr CLI. Prefers an already-installed binary, then the latest GitHub release, then a build from source with go install.
license: MIT
compatibility: Requires curl (or wget) and tar to install from a release, or a Go toolchain for the source fallback. Standalone — does not need the vultr CLI to be present already. Installs from the public repository github.com/ideamans/vultr-cli, so no GitHub authentication is needed.
allowed-tools: Bash(curl:*) Bash(wget:*) Bash(tar:*) Bash(unzip:*) Bash(go:*) Bash(uname:*) Bash(command:*) Bash(which:*) Bash(mkdir:*) Bash(mv:*) Bash(cp:*) Bash(rm:*) Bash(chmod:*) Bash(ls:*) Bash(test:*) Bash(echo:*) Read
---

# vultr-cli-install

Make the `vultr` command usable, doing the least work that achieves it.

## Route 1 — an existing installation on PATH

```bash
command -v vultr && vultr --version
```

If that resolves, **use it and stop here.** Do not check for a newer release —
it costs an API call and the user did not ask for an upgrade.

Two checks before trusting the hit:

- **It is the right tool.** Vultr Inc. publishes its own, unrelated CLI, and
  other things may own the name. `vultr --version` must report a version string;
  `vultr llm | head -1` must read `# vultr CLI — reference for AI agents`. If
  something else owns the name, tell the user and use an explicit path to this
  binary rather than shadowing theirs.
- **It is recent enough.** If `vultr llm` is not a known command, the binary
  predates the embedded reference. Say so and continue to route 2 to upgrade it.

Continue past this section only when the command is missing, is the wrong tool,
is too old, or the user explicitly asked to update.

## Route 2 — the latest GitHub release

The repository is public, so no authentication is needed.

```bash
VERSION=$(curl -fsSL https://api.github.com/repos/ideamans/vultr-cli/releases/latest \
  | grep '"tag_name"' | head -1 | cut -d'"' -f4)   # e.g. v0.2.0
```

Archive names follow goreleaser's defaults for this project:

```
vultr-cli_<version-without-v>_<os>_<arch>.tar.gz
```

`<os>` is `darwin`, `linux` or `windows` (lowercase) and `<arch>` is `amd64` or
`arm64` — so `uname -m` reporting `x86_64` maps to `amd64`. Windows ships a
`.zip` rather than a `.tar.gz`.

```bash
OS=$(uname -s | tr '[:upper:]' '[:lower:]')            # darwin | linux
ARCH=$(uname -m); [ "$ARCH" = "x86_64" ] && ARCH=amd64  # amd64 | arm64
curl -fsSL -o /tmp/vultr-cli.tar.gz \
  "https://github.com/ideamans/vultr-cli/releases/download/${VERSION}/vultr-cli_${VERSION#v}_${OS}_${ARCH}.tar.gz"
```

If the download 404s, list the actual assets on the release page rather than
retrying variations — the naming is set by `.goreleaser.yaml` and may have
changed.

### Install onto PATH

The binary inside the archive is named `vultr`, not `vultr-cli`.

```bash
tar -xzf /tmp/vultr-cli.tar.gz -C /tmp
mkdir -p ~/.local/bin && mv /tmp/vultr ~/.local/bin/ && chmod +x ~/.local/bin/vultr
```

Prefer the first writable directory already on PATH — `~/.local/bin`, then
`/usr/local/bin`. Two things not to do on your own initiative:

- If nothing on PATH is writable, leave the binary in `/tmp`, print the exact
  `sudo mv` command and let the user run it. Do not run `sudo` yourself.
- If `~/.local/bin` is not on PATH, give the user the line to add to their shell
  profile. Do not edit the profile for them.

## Route 3 — build from source

Needs a Go toolchain and compiles rather than downloads, so it is the last
resort — but it covers platforms the release assets miss.

```bash
go install github.com/ideamans/vultr-cli@latest
```

The binary lands in `$(go env GOPATH)/bin` and is named `vultr-cli` here, not
`vultr`, because `go install` uses the module name. Either add a symlink or tell
the user which name to invoke — do not leave them guessing why `vultr` is still
missing.

## Verify

```bash
vultr --version
vultr llm | head -5
```

Report which route was taken, the version and the install path — the user needs
to know whether anything was written to their machine. Then continue with what
they originally asked for.

Nothing works until `VULTR_API_KEY` is set, so if the user has not configured it
yet, say so now rather than letting the first real command fail.
