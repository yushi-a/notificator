# notificator

通知リクエストを受け取り、LINE にプッシュ送信する Go 製のサーバーです。API 定義は別リポジトリ [yuxsr-dev-pb](https://github.com/yushi-a/yuxsr-dev-pb) の `NotificatorService` を使っています。

## 概要

- Connect / gRPC / gRPC-Web を同一ポート（`:50051`）で受け付ける（HTTP/1.1 と TLS なしの HTTP/2 に対応）
- 受け取ったメッセージを LINE Messaging API（line-bot-sdk-go）でプッシュ送信する
- CLI は Cobra。`serve` サブコマンドでサーバーを起動する

## 構成

| パス | 内容 |
| --- | --- |
| `main.go` | エントリーポイント |
| `cmd/` | Cobra コマンド（`root.go`、`serve.go`）と設定の読み込み |
| `internal/server/` | HTTP サーバーとハンドラ登録 |
| `internal/server/service/` | `NotificatorService` の実装とクライアントのインターフェース |
| `internal/client/` | LINE クライアント |
| `internal/logging/` | ロガー |
| `Dockerfile` | マルチステージのイメージ（実行は `app serve`） |

## 必要なもの

- Go（`go.mod` は 1.27.3）
- LINE の設定値を環境変数で渡す: `LINE_CHANNEL_SECRET` / `LINE_CHANNEL_ACCESS_TOKEN` / `LINE_USER_ID`
  - `$HOME/.notificator.yaml` または `--config` で指定する設定ファイルからも読める
- `golangci-lint`、`govulncheck`（lint / 脆弱性チェック用）
- Docker（イメージをビルドするとき）

## 動かし方

```bash
make run        # go run main.go（引数なしだとヘルプ。サーバー起動は go run main.go serve）
make lint       # golangci-lint run ./...
make vuln       # govulncheck ./...
make build      # 静的バイナリ app を生成
go test -race ./...
```

## デプロイ

`.github/workflows/ci.yaml` が lint・テスト・govulncheck を通したあと Docker イメージをビルドします。`v*` タグのときだけ `ghcr.io/yushi-a/notificator` に push します。CI ではダミーの LINE 設定で `serve` を起動して応答を確認し、Trivy でスキャンします。クラスタへのデプロイ定義（helm など）はこのリポジトリにはありません（TODO: 別リポジトリで管理しているか確認）。

## ライセンス

TODO: `LICENSE` ファイルはありますが中身が空です。ライセンスを決めて記入してください。
