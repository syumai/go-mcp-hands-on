# mcp-goを使ったMCPサーバーのテスト

mcp-goを使ったMCPサーバーのテストは、以下を組み合わせることで簡単に行えます。

* mcp-goのSSEClient
* mcp-goのSSEServer
* httptest package

標準入出力を使ったテストも可能だとは思いますが、Goには標準で利用可能なHTTPサーバーのモックに使えるhttptest packageがあるので、こちらを使うのが楽でしょう。

## テストの実装

`06/caesar-mcp/server/server_test.go` の内容を上から解説していきます。

## 課題

テストケースが一つ空欄で存在しているので、実装してください。
