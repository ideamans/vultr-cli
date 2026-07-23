---
name: vultr-cli-usage
description: Operate Vultr cloud infrastructure from the command line with the vultr CLI — instances, bare metal, block storage, DNS, load balancers, Kubernetes, firewalls, snapshots and every other Vultr API service. Use when the user asks to list, create, inspect, resize or delete Vultr resources, mentions their Vultr account, or is working with a VPS they host there.
license: MIT
compatibility: Requires the `vultr` binary on PATH — run the vultr-cli-install skill if it is missing. Needs VULTR_API_KEY in the environment or an --api-key value. Every command hits the live Vultr API; there is no dry-run mode.
allowed-tools: Bash(vultr:*) Bash(jq:*) Bash(command:*) Read Write
---

# vultr-cli-usage

Drive the Vultr API through the `vultr` CLI. It covers every operation govultr
exposes, which is far more than any one session needs — so the workflow is
always *look it up, then run it*.

## 1. Confirm the tool and the credentials

```bash
command -v vultr && vultr --version
```

Missing? Run the `vultr-cli-install` skill, then come back.

```bash
vultr account get
```

This is the cheapest call that proves the key works. `Error: no API key: set
VULTR_API_KEY or pass --api-key` means the environment is not configured — ask
the user to export `VULTR_API_KEY`; do not put the key on the command line where
it lands in shell history.

If requests fail intermittently with 401, the key is probably IP-restricted and
resolving over IPv6. Add `-4`.

## 2. Find the operation

Do **not** guess command names. The reference is embedded in the binary:

```bash
vultr llm | grep -i 'vultr instance'
vultr llm | less                       # whole reference, ~800 lines
```

Names are kebab-case renderings of the govultr methods: `Instance.List` →
`vultr instance list`, `Database.ListConnectionPools` →
`vultr database list-connection-pools`.

## 3. Compose the call

Read the signature in the catalog before writing anything.

- `<json:TypeName>` takes a JSON body: inline, `@file.json`, or `-` for stdin.
  **Unknown fields are rejected**, so check the shape first:

  ```bash
  vultr instance create --schema
  ```

  `--schema` prints the accepted fields and exits without calling the API.
- `*ListOptions` in the signature means pagination comes from *flags*
  (`--per-page`, `--cursor`, and where supported `--tag`, `--label`, `--region`,
  `--main-ip`, `--description`) — not from a positional argument.
- Some list operations take a required filter string first. Pass `''` for "all":
  `vultr plan list '' --per-page 5`.
- To enumerate everything, use `--all` rather than looping on `--cursor`.

## 4. Run it, and treat destructive commands as destructive

Output is JSON on stdout, so pipe it into `jq` directly:

```bash
vultr instance list --all | jq '.data[] | {id, label, main_ip}'
```

**There is no confirmation prompt and no dry run.** `vultr instance delete <id>`
deletes immediately, and there is no undo for most resources. Before running any
delete, destroy, detach or halt operation: state in prose exactly what will be
affected, and get the user's agreement. When the target came from a list, quote
the label and IP alongside the ID so they can recognise it.

## 5. Report

Give back the identifier that was created or changed, and the fields the user
will need next (IP address, status, region). On failure, quote the `Error:` line
verbatim — the CLI's messages are actionable as written.

## Failure modes

| Symptom | Cause | Fix |
| --- | --- | --- |
| `command not found: vultr` | not installed | run the `vultr-cli-install` skill |
| `Error: no API key: ...` | `VULTR_API_KEY` unset | ask the user to export it |
| intermittent 401 | IP-restricted key over IPv6 | add `-4` |
| `Error: json: unknown field "..."` | body field guessed | re-run with `--schema` and match it exactly |
| `Error: too many arguments: ...` | passed a positional where flags are expected | the signature has `*ListOptions` — use `--per-page` etc. |
| exit 0 with no output | the operation returns nothing | expected; confirm with the matching `get` or `list` |
