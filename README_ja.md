# vultr-cli

[English](README.md) | 日本語

Vultr API 全体をラップする aws 風 CLI です。[govultr](https://github.com/vultr/govultr)（Vultr 公式 Go クライアント）の全サービス・全メソッドをリフレクションで自動的にサブコマンド化しているため、Vultr が提供する全 API（35 サービス・428 オペレーション）をカバーします。govultr は Vultr 公式 CLI（vultr/vultr-cli）も依存する公式リファレンスクライアントです。

## ビルド

```bash
go build -o vultr .
```

## 認証

環境変数 `VULTR_API_KEY` に API キーを設定するか、`--api-key` フラグで指定します。

> **注意**: Vultr の API キーには「Allowed IPs」の制限を設定できます。`{"error":"Unauthorized IP address: ..."}` が返る場合は、Vultr コンソール（Account → API）で現在の IP アドレスを許可してください。
>
> IPv4 のみ許可している環境では、接続が IPv4/IPv6 で交互に切り替わり**断続的に 401 になる**ことがあります。その場合は `-4` / `--ipv4` フラグで IPv4 接続を強制してください。IPv6 も併用する場合は、プライバシー拡張でアドレス下位 64 ビットが変動するため `/64` プレフィックス単位で許可するのが確実です。
>
> ```bash
> vultr -4 instance list   # IPv4 を強制
> ```

## 使い方

```
vultr <service> <operation> [args...]
```

サービス・オペレーション一覧は `--help` で確認できます。

```bash
vultr --help                # 全サービス一覧
vultr instance --help       # instance の全オペレーションとシグネチャ
vultr instance get --help   # 個別オペレーションの引数
```

### AIエージェントから使う

`vultr llm` は AI エージェントが読むことを想定した自己完結のリファレンス（利用規約・全オペレーションのカタログ・全リクエストボディの JSON スキーマ）を一括出力します。カタログとスキーマはコマンドツリーと同じリフレクションデータから生成するため実装と乖離せず、バイナリに埋め込まれているのでオフラインでも実行中のバージョンと必ず一致します。

```bash
vultr llm                  # Markdown
vultr llm --format json    # 章ごとの JSON 配列
vultr --llm                # 非推奨エイリアス。従来どおりどの位置でも動作します
```

Claude Code ではプラグインを導入すると `/vultr-usage` と `/vultr-install` が使えます。

```
/plugin marketplace add ideamans/claude-public-plugins
/plugin install vultr-cli@ideamans-plugins
```

同じスキルは Copilot や Cursor など Agent Skills 対応ホストでも利用できます。

```bash
gh skill install ideamans/vultr-cli/plugins/vultr-cli/skills/vultr-usage --agent copilot
```

スキル本体は [`plugins/vultr-cli/`](plugins/vultr-cli)、準拠している標準は [ideamans/go-llm-cli-kit](https://github.com/ideamans/go-llm-cli-kit) を参照してください。

### 例

```bash
# アカウント情報
vultr account get

# インスタンス一覧（ページング）
vultr instance list --per-page 100
vultr instance list --all            # 全ページを自動取得

# インスタンス取得・操作
vultr instance get <instance-id>
vultr instance start <instance-id>
vultr instance reboot <instance-id>
vultr instance delete <instance-id>

# インスタンス作成（JSON ボディはインライン / @ファイル / stdin）
vultr instance create '{"region":"nrt","plan":"vc2-1c-1gb","os_id":2136,"label":"test"}'
vultr instance create @body.json
cat body.json | vultr instance create -

# リクエストボディの構造を確認
vultr instance create --schema

# DNS
vultr domain list
vultr domain-record list example.com
vultr domain-record create example.com '{"name":"www","type":"A","data":"192.0.2.1","ttl":300}'

# 複数 ID を取る操作はカンマ区切り
vultr instance mass-reboot id1,id2,id3

# その他のサービスも同じパターン
vultr kubernetes list-clusters
vultr block-storage list
vultr firewall-group list
```

### 引数の規則

| シグネチャの型 | 渡し方 |
|---|---|
| `<string>` `<int>` `<bool>` | 位置引数 |
| `<json:XxxReq>` | JSON 文字列 / `@file.json` / `-`（stdin）。`--schema` で構造を確認 |
| `<v1,v2,...>`（`[]string`） | カンマ区切り or JSON 配列 |
| `[<bool\|null>]` などの省略可能引数 | 省略すると `null`、明示的に `null` も指定可 |
| `*ListOptions` | `--per-page` `--cursor` `--tag` `--label` `--region` などのフラグ |

### 出力

- 結果は JSON（2 スペースインデント）で標準出力に出力
- リスト系はページネーション情報付き: `{"data": [...], "meta": {...}}`
- `--all` を付けるとカーソルを自動追跡して全件を結合
- 戻り値のない操作（delete / start など）は成功時に何も出力せず終了コード 0
- エラーは標準エラー出力に出力し終了コード 1

## 仕組み

`govultr.Client` の各サービスフィールド（インターフェース）をリフレクションで走査し、コマンドツリーを自動生成しています。API の実装は govultr に完全に委譲しているため、`go.mod` の govultr を更新して再ビルドするだけで新 API に追従します。`dispatch_test.go` がディスパッチャの前提とするシグネチャ不変条件を全 428 メソッドに対して検証しており、govultr の更新で前提が崩れた場合は実行時の誤動作ではなくテストの失敗として検出されます。

- `main.go` — エントリポイント・認証・JSON 出力
- `commands.go` — リフレクションによるコマンド生成とディスパッチ
- `args.go` — 位置引数の型変換（スカラー / JSON / @file / stdin）と `--schema` 生成
- `naming.go` — Go 識別子 → kebab-case 変換（`CreateIPv4` → `create-ipv4` 等）
- `llm.go` — リクエストボディ JSON スキーマ章の生成
- `gen_llmdocs.go` — `go generate` で埋め込みリファレンスを再生成（隠しコマンド `gen-llmdocs`）
- `internal/llmdocs/` — `vultr llm` が出力する埋め込みリファレンス（`00-guide.md` は手書き、`90-`/`91-` は生成物）
- `dispatch_test.go` — ディスパッチャ前提条件の全数監査テスト

## CI / リリース

- GitHub Actions が push / PR で vet / build / `go test -race` を実行します（`.github/workflows/test.yml`）
- `v*` タグの push で [GoReleaser](https://goreleaser.com/) がクロスプラットフォームバイナリ（linux / darwin / windows × amd64 / arm64）を GitHub Releases に公開します（`.github/workflows/release.yml`、`.goreleaser.yaml`）

```bash
git tag v0.1.0 && git push origin v0.1.0
```
