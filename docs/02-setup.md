# 環境構築

## MCPホストのセットアップ

MCPサーバーを開発して動かすには、MCPサーバーと接続するクライアントを管理するホストが必要となるため、何かしらのMCPホストを事前にインストールしておいてください。

現在、非常にハードルの低いMCPホストとしておすすめなのは、[VS Code (GitHub Copilot)](https://code.visualstudio.com/)です。
Copilot FreeでもMCP Serverの登録が可能です。
https://code.visualstudio.com/docs/copilot/chat/mcp-servers

筆者自身が最もよく使っているAIエージェントは[Claude Desktop](https://claude.ai/download)です。
独立したアプリケーションとして使えますし、MCPサーバーとの連携がこなれていて使いやすいのでおすすめです。

MCPサーバーの設定は以下のリンクを参考に行えます。
https://modelcontextprotocol.io/quickstart/user

本ハンズオンではVS Code (GitHub Copilot)、あるいはClaude Desktopの使用を前提としますが、以下のツールに関してもサポート可能です。

### VS Code拡張系

* [Cline](https://github.com/cline/cline)
* [Roo Code](https://github.com/RooVetGit/Roo-Code)

### CLI tool系

* [Claude Code](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview)
  - https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/tutorials#set-up-model-context-protocol-mcp
* [oterm](https://github.com/ggozad/oterm)

以下のエディターでもMCPサーバーとの連携が可能ですが、筆者はまだ設定を行ったことがないため、ご自身でドキュメントを参考に作業を行ってください。　

* [Cursor](https://github.com/cursor-ai/cursor)
  - https://docs.cursor.com/context/model-context-protocol
* [Windsurf](https://github.com/windsurf-io/windsurf)
  - https://docs.windsurf.com/windsurf/mcp

## 開発環境のセットアップ

Goさえインストールしてあれば、エディタは何でも構いません。ご自身の好きなエディタをご利用ください。
