# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

これはGo言語でMCP（Model Context Protocol）サーバーを実装するハンズオン資料のリポジトリです。caesar-mcpディレクトリに実装されたシーザー暗号のMCPサーバーが含まれています。

## アーキテクチャ

- **caesar-mcp/**: メインのMCPサーバー実装
  - `main.go`: MCPサーバーのエントリーポイント
  - `server/server.go`: MCPサーバーの定義とツール登録
  - `caesar/caesar.go`: シーザー暗号のコア実装（RotN関数）
  - `server/server_test.go`: MCPサーバーのテスト
- **docs/**: ハンズオン資料（Markdown形式）
- **examples/**: サンプルコード

## 重要なコマンド

### ビルドとテスト
```bash
# caesar-mcpディレクトリに移動
cd caesar-mcp

# プロジェクトをビルド
go build

# テスト実行
go test ./...

# MCPサーバーを実行（stdioモード）
go run main.go
```

## MCP実装の構造

このプロジェクトでは `github.com/mark3labs/mcp-go` ライブラリを使用してMCPサーバーを実装しています：

1. `server.NewMCPServer()` でサーバーインスタンスを作成
2. サーバーにツール（caesar_rotate）を登録
3. `server.ServeStdio()` でstdio経由でMCPプロトコルを処理

現在の実装では、server.goのrotateHandlerが未実装の状態で、これがハンズオンの実装対象となっています。

## テスト環境

- Caesar暗号の機能テスト: `caesar/caesar_test.go`
- MCPサーバーのツールテスト: `server/server_test.go`

テストの一部が失敗する場合は、MCPツールの登録が未完了の可能性があります。