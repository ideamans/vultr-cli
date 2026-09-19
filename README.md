# vultr-cli

English | [日本語](README_ja.md)

An aws-style CLI that wraps the entire Vultr API. It reflects over every service and method of [govultr](https://github.com/vultr/govultr) (Vultr's official Go client) and turns them into subcommands automatically, covering the full Vultr API surface (35 services, 428 operations). govultr is the official reference client that Vultr's own CLI (vultr/vultr-cli) depends on as well.

## Install

Download the archive for your platform from [Releases](https://github.com/ideamans/vultr-cli/releases) and put the `vultr` binary on your `PATH`.

To build from source:

```bash
go build -o vultr .
```

Claude Code users can let the `/vultr-install` skill do this (see below).

## Authentication

Set your API key in the `VULTR_API_KEY` environment variable, or pass it with the `--api-key` flag.

> **Note**: Vultr API keys can be restricted with an "Allowed IPs" list. If you get `{"error":"Unauthorized IP address: ..."}`, allow your current IP address in the Vultr console (Account → API).
>
> When only IPv4 addresses are allowed, connections may alternate between IPv4 and IPv6 and **fail intermittently with 401**. In that case, force IPv4 with the `-4` / `--ipv4` flag. If you also want to use IPv6, allow the `/64` prefix rather than a single address, since IPv6 privacy extensions rotate the lower 64 bits.
>
> ```bash
> vultr -4 instance list   # force IPv4
> ```

## Usage

```
vultr <service> <operation> [args...]
```

Discover services and operations with `--help`:

```bash
vultr --help                # list all services
vultr instance --help       # all instance operations with signatures
vultr instance get --help   # arguments of a single operation
```

### Use from an AI agent

`vultr llm` prints a single self-contained reference: usage conventions, the full catalog of all operations, and the JSON schema of every request body type. The catalog and schemas are generated from the same reflection data as the command tree, so they never drift from the actual CLI, and the whole thing is embedded in the binary — it works offline and always matches the version you are running.

```bash
vultr llm                  # Markdown
vultr llm --format json    # chapters as a JSON array
vultr --llm                # deprecated alias, still accepted anywhere on the line
```

Claude Code users can install the plugin instead, which adds `/vultr-usage` and `/vultr-install`:

```
/plugin marketplace add ideamans/claude-public-plugins
/plugin install vultr-cli@ideamans-plugins
```

The same skills work in Copilot, Cursor and other Agent Skills hosts:

```bash
gh skill install ideamans/vultr-cli/plugins/vultr-cli/skills/vultr-usage --agent copilot
```

See [`plugins/vultr-cli/`](plugins/vultr-cli) for the skills themselves, and [ideamans/go-llm-cli-kit](https://github.com/ideamans/go-llm-cli-kit) for the standard they follow.

### Examples

```bash
# Account info
vultr account get

# List instances (pagination)
vultr instance list --per-page 100
vultr instance list --all            # fetch all pages automatically

# Get / operate on an instance
vultr instance get <instance-id>
vultr instance start <instance-id>
vultr instance reboot <instance-id>
vultr instance delete <instance-id>

# Create an instance (JSON body: inline / @file / stdin)
vultr instance create '{"region":"nrt","plan":"vc2-1c-1gb","os_id":2136,"label":"test"}'
vultr instance create @body.json
cat body.json | vultr instance create -

# Inspect the expected request body structure
vultr instance create --schema

# DNS
vultr domain list
vultr domain-record list example.com
vultr domain-record create example.com '{"name":"www","type":"A","data":"192.0.2.1","ttl":300}'

# Operations taking multiple IDs use comma-separated values
vultr instance mass-reboot id1,id2,id3

# Every other service follows the same pattern
vultr kubernetes list-clusters
vultr block-storage list
vultr firewall-group list
```

### Argument rules

| Type in signature | How to pass |
|---|---|
| `<string>` `<int>` `<bool>` | positional argument |
| `<json:XxxReq>` | JSON string / `@file.json` / `-` (stdin); inspect with `--schema` |
| `<v1,v2,...>` (`[]string`) | comma-separated values or a JSON array |
| optional args like `[<bool\|null>]` | omit for `null`, or pass `null` explicitly |
| `*ListOptions` | flags such as `--per-page` `--cursor` `--tag` `--label` `--region` |

### Output

- Results are printed to stdout as JSON (2-space indent)
- List operations include pagination info: `{"data": [...], "meta": {...}}`
- With `--all`, the cursor is followed automatically and all pages are merged
- Operations with no return value (delete / start etc.) print nothing and exit 0 on success
- Errors go to stderr with exit code 1

## How it works

The CLI walks every service field (interface) of `govultr.Client` via reflection and generates the command tree. All API implementation is delegated to govultr, so tracking new APIs only requires bumping govultr in `go.mod` and rebuilding. `dispatch_test.go` verifies the signature invariants the dispatcher relies on across all 428 methods, so a govultr upgrade that breaks an assumption fails the tests instead of misbehaving at runtime.

- `main.go` — entry point, authentication, JSON output
- `commands.go` — reflection-based command generation and dispatch
- `args.go` — positional argument conversion (scalars / JSON / @file / stdin) and `--schema` generation
- `naming.go` — Go identifier → kebab-case conversion (`CreateIPv4` → `create-ipv4` etc.)
- `llm.go` — request body JSON schema chapter generation
- `gen_llmdocs.go` — regenerates the embedded reference under `go generate` (hidden `gen-llmdocs` command)
- `internal/llmdocs/` — the reference `vultr llm` prints (`00-guide.md` hand-written, `90-`/`91-` generated)
- `dispatch_test.go` — full-surface audit of dispatcher assumptions

## License

MIT

## CI / Release

- GitHub Actions runs vet / build / `go test -race` on pushes and pull requests (`.github/workflows/test.yml`).
- Pushing a `v*` tag triggers [GoReleaser](https://goreleaser.com/) to publish cross-platform binaries (linux / darwin / windows, amd64 / arm64) to GitHub Releases (`.github/workflows/release.yml`, `.goreleaser.yaml`).

```bash
git tag v0.1.0 && git push origin v0.1.0
```
