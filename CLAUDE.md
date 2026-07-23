# CLAUDE.md — vultr-cli

Vultr API 全体をラップする aws 風 CLI。`govultr` の全サービス・全メソッドを
リフレクションでサブコマンド化している（35サービス / 428オペレーション）。
**バイナリ名は `vultr`**（リポジトリ名と異なる。`go install` した場合のみ
`vultr-cli` になる）。

## 変更時の必須手順

**機能を追加した、フラグを増やした、既存の挙動を変えた — このいずれかをしたら、
3か所すべてを更新してから終わること。**

| 更新先 | 対象 | やり方 |
| --- | --- | --- |
| ① ドキュメント | `README.md` / `README_ja.md` | 手で更新。使い方が変わったときのみ |
| ② ヘルプ | cobra の `Short` / `Long` / フラグ説明 | コード内 |
| ③ **LLMナレッジ** | `internal/llmdocs/00-guide.md` | 手書き。鉄則・認証・ワークフロー・失敗モードが変わったら |
| | `internal/llmdocs/90-commands.md`<br>`internal/llmdocs/91-schemas.md` | **生成物。手編集しない** → `go generate ./...` |
| | `plugins/vultr-cli/skills/*/SKILL.md` | 手順や前提が変わったとき |
| | `context7.json` の `rules` | 新しい落とし穴が生まれたとき |

③ を忘れやすい。ドキュメントとヘルプは人間が読んで気づくが、**LLMナレッジが
古いことには誰も気づかない**（エージェントが黙って間違えるだけ）。

判断に迷ったときの目安:

- **govultr を更新した** → `go generate ./...` は必須。新サービス・新オペレーションが
  カタログに入る。シグネチャの前提が崩れていれば `dispatch_test.go` が落ちる
- 新しいフラグを足した → ② と `go generate`。使い方が非自明なら `00-guide.md` にも
- 出力形式・終了コードを変えた → ①②③すべて。特に `00-guide.md` の「Output contract」と
  `context7.json` の該当 rule
- エージェントが間違えやすい罠を見つけた → `00-guide.md` の失敗モード表と
  `context7.json` の `rules`
- 破壊的操作の扱いを変えた → `vultr-cli-usage` の SKILL.md。この CLI には確認プロンプトが
  無いため、スキル側で同意を取る手順を保っていること

## リリース

`pluginVersion`（`main.go`）と `plugin.json` の `version` と git タグの3つを揃える。
テストとリリースワークフローが不一致を検出する。手順は
`plugins/vultr-cli/PUBLISH.md`。

## 確認

```bash
go generate ./...     # 生成物を作り直す
git diff --exit-code  # 差分が出たらコミット漏れ
go test ./...         # SKILL.md 検証とバージョン整合を含む
go run . llm | head   # 埋め込みリファレンスが壊れていないか
```

## 参照

- 標準: <https://github.com/ideamans/go-llm-cli-kit/blob/main/LLM.md>
- 生成物と原本の対応: `.claude/rules/ai-artifacts-policy.md`
- 再生成: `/regen-ai`
