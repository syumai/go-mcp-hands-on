# mcp-goを使ったMCPサーバーの実装

## お題

シーザー暗号を解くためのMCPサーバーを実装します。コマンド名は、 `caesar-mcp` です。

シーザー暗号は、非常にシンプルな暗号で、単に渡された文字列を、指定された数字の分、順番にアルファベットをずらすことで暗号化します。

復号は、ずらした分を逆方向にずらすことで行います。もしくは、アルファベットが全部で26文字なので、26からずらした分を引き、その分を同じ方向にずらすことでも復号できます。

## シーザー暗号の例

暗号化: `Hello` を5文字ずらす -> `Mjqqt`
復号: `Mjqqt` を5文字逆方向にずらす -> `Hello`

## シーザー暗号をお題にする理由

生成AIは、シーザー暗号の「アルファベットをずらす」といった処理や、計算を苦手とします。
質問に対して、誤った答えを返す可能性が非常に高いです。

例えば、次の質問をご自身の生成AIに投げてみてください。

```
このシーザー暗号を解いてください。
「Xf nrj uvjzxevu sp Ifsvik Xizvjvdvi, Ifs Gzbv, reu Bve Kyfdgjfe.」
```

この問題の答えは、

「Go was designed by Robert Griesemer, Rob Pike, and Ken Thompson.」

です。
正しい答えが返ってきたでしょうか？
筆者の場合、Claude 3.7 Sonnetは誤った回答を行いました。
o4-miniやo3のような、回答にプログラムを使用する生成AIエージェントは正しい回答を行えてしまうので、それ以外で試してみてください。

## 課題

シーザー暗号を解くMCPサーバーを実装し、生成AIに以下の暗号文を解かせてください。

1. Twnts nruqjrjsyji fs fxdshmwtstzx rniiqjbfwj yt nsyjwhjuy vzfyjwsfwd FUN wjvzjxyx.
2. Lmdm dqrmofadqp ftq baxkyadbtuo uztqdufmzoq efdgofgdq rad nqffqd qzombegxmfuaz.
3. Erltgwxk fxmbvnehnler wxunzzxw max kxvnklbox tezhkbmaf vtnlbgz lmtvd hoxkyehp xqvximbhgl.

## 実装について

実装に使う `RotN` 関数は、 `05/caesar-mcp/caesar/caesar.go` に実装されています。
この関数を使って、`05/caesar-mcp/server/server.go` を実装してください。

ツールのパラメーターは、以下の2つです。

* text: 文字列
* shift: 数値
  - デフォルト値: 13

## MCPサーバーの登録方法

VS Code (GitHub Copilot) をご利用の場合、以下の `${userHome}` のような事前定義変数を使って、ユーザーのホームディレクトリのパスを省略表記できます。ワークスペースのパスを利用する変数も存在します。

https://code.visualstudio.com/docs/reference/variables-reference

VS Code以外をご利用の方は、フルパスをご利用ください。

```json
"servers": {
  "caesar-mcp": {
    "type": "stdio",
    "command": "go",
    "args": [
      "run",
      "${userHome}/go/src/github.com/syumai/go-mcp-hands-on/05/caesar-mcp"
    ]
  }
}
```