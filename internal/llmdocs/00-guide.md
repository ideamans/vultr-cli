# vultr CLI — reference for AI agents

An aws-style CLI wrapping the entire Vultr API through
[govultr](https://github.com/vultr/govultr). Every API operation is reachable;
nothing prompts for input, so it is safe to run from scripts and agents.

```
vultr <service> <operation> [args...] [flags]
```

This reference is embedded in the binary — `vultr llm` always describes the
exact version you are running. The command catalog and request body schemas
below are generated from the same reflection data the CLI dispatches on, so they
cannot drift from what the binary actually accepts.

## Authentication

Set `VULTR_API_KEY` in the environment, or pass `--api-key <key>`. Without
either, every command fails with `no API key: set VULTR_API_KEY or pass
--api-key`.

If the API key is IP-restricted and requests fail intermittently with 401, add
`-4` (`--ipv4`) to force IPv4 connections.

## Argument conventions

Arguments are positional and typed by the operation signature shown in the
catalog:

| Placeholder | How to pass it |
| --- | --- |
| `<string>` `<int>` `<bool>` | a plain positional value |
| `<json:TypeName>` | a JSON request body — inline JSON, `@path/to/file.json`, or `-` to read stdin |
| `<v1,v2,...>` | a `[]string` — comma-separated values or a JSON array |
| `[<bool\|null>]` | optional trailing argument — omit it, or pass the literal `null` |

Rules that matter:

- **Unknown JSON fields are rejected.** Do not guess field names. Every
  `TypeName` is listed in the request body schemas chapter, and
  `vultr <service> <operation> --schema` prints the schema for that one
  operation and exits without calling the API. Use it to check a body before
  sending it.
- **Some list operations take a required string filter first.** Pass an empty
  string `''` for "all" — e.g. `vultr plan list '' --per-page 5`.
- Operations whose signature includes `*ListOptions` take pagination and
  filtering *flags* rather than a positional argument: `--per-page <n>`,
  `--cursor <c>`, and where the API supports them `--tag`, `--label`,
  `--region`, `--main-ip`, `--description`.
- List operations that return `Meta` also accept `--all`, which follows
  pagination and merges every page into one result. Prefer it over looping on
  `--cursor` yourself.

## Output contract

- Output is JSON on **stdout**. A single return value prints as-is.
- Multiple return values print as an object: `{"data": ..., "meta": ...}`.
  Extra values are keyed by their snake_case type name, e.g.
  `vultr database list-connection-pools` →
  `{"data": ..., "database_connection_pools": [...], "meta": {...}}`.
- Operations with no return value print nothing and exit 0.
- Errors go to **stderr** prefixed with `Error:`; the exit code is 1.

Because the output is already JSON, pipe it straight into `jq` rather than
parsing prose.

## Typical workflows

### Find the operation you need

The catalog is large — several hundred operations across every Vultr service.
Search this reference rather than guessing a command name:

```bash
vultr llm | grep -i 'vultr instance'
vultr llm --format json | jq '.[] | select(.title == "Command catalog")'
```

Service and operation names are kebab-case renderings of the govultr method
names: `Instance.List` → `vultr instance list`,
`Database.ListConnectionPools` → `vultr database list-connection-pools`.

### Create a resource from a JSON body

```bash
vultr instance create --schema              # inspect the accepted fields first
vultr instance create '{"region":"nrt","plan":"vc2-1c-1gb","os_id":1743}'
vultr instance create @instance.json        # or from a file
```

### Enumerate everything

```bash
vultr instance list --all | jq '.data[].id'
```

## Failure modes

| Symptom | Cause | Fix |
| --- | --- | --- |
| `Error: no API key: ...` | credentials missing | set `VULTR_API_KEY` or pass `--api-key` |
| intermittent 401 | IP-restricted key resolving over IPv6 | add `-4` |
| `Error: json: unknown field "..."` | guessed a body field name | run the operation with `--schema` and match it |
| `Error: too many arguments: ...` | passed a value where the operation takes a flag | check the signature in the catalog — `*ListOptions` means flags, not a positional |
| empty output, exit 0 | the operation returns nothing | expected; verify with a corresponding `get`/`list` |

## What this CLI will not do

- It does not retry, rate-limit or back off. A 429 surfaces as an error.
- It does not maintain state between invocations — no profiles, no cached IDs.
  Carry identifiers forward from earlier JSON output.
- It does not confirm destructive operations. `vultr instance delete <id>`
  deletes immediately. Confirm with the user in prose before running one.
