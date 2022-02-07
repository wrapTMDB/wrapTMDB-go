
<h3 align="right">
<a href="https://github.com/wrapTMDB/wraptmdb-go">Github page</a> |
<a href="https://pkg.go.dev/github.com/wrapTMDB/wraptmdb-go">pkg.go</a>  
</h3>


# WrapTMDB-ts  
<h3>
<p align="center">
<a href="README.md"> English </a>|
<a href="/docs/README_ja.md"> 日本語 </a>|
<a href="/docs/README_zh-tw.md"> 繁體中文 </a>|
<a href="/docs/README_zh-ch.md"> 简体中文 </a>
</p>
</h3>
<br/>

# [wrapTMDB](https://github.com/wrapTMDB/wrapTMDB)とは何ですか ?

```wrapTMDB``` は,TMDB APIをラップし,ドキュメントから,さまざまなプログラム言語で実装するのラッパーコレクションです。

開発者が情報やメタデータについて映画やテレビ番組をリクエストするのに役立ちます。<br/>

Golangに作て, pkg.goで公開されるリポジトリです,<br/>
続きを [見る](https://github.com/wrapTMDB/wrapTMDB).
___
## どのようなプロジェクトがこのツールを使用しますか ?

- 新しい映画情報を追跡するクライアントを作成したい場合。
- ムービーファイルやビデオの管理に役立つツールを作成したい場合。.
- TMDB公式サイトに取り替えたい場合でも。 (( www
- ...

___
## Useage

### Install:

```bash
$go get github.com/wrapTMDB/wraptmdb-go
```

このツールを使用する前に, [api_key](https://developers.themoviedb.org/3/getting-started/authentication) がすでにあることを確認してください.
<br/>

``` Golang
package main

import (
	"fmt"
	wraptmdb "github.com/wrapTMDB/wraptmdb-go"
)

func main() {
	//initialize
	wraptmdb.Init("Your api_key")
	wraptmdb.SetHeader(map[string]string{
		"User-Agent": "wraptmdb-go dev",
		"Referer":    "wraptmdb-go",
	})
	//call function
	msg := wraptmdb.Movies.GetDetail("624860", "")

	fmt.Print(msg)
}

```
___

## これらのAPIをどのように認識しますか ?

### 直感を使う:
```Golang
data := wrapTMDB.Movies.GetDetails("624860");
```
![alt text](172714.png)

```Golang
data := wrapTMDB.Collections.GetTranslations("654321", "en-US");
```
![alt text](172927.png)

```Golang
data := wrapTMDB.TVseasons.GetImages("54321", "65421", "en-US");
```
![alt text](172331.png)



# 開発に参加する ?
```bash
$git clone https://github.com/wrapTMDB/wraptmdb-go &&

npm install ||

touch src/main.go
```

___
## その他

*** 星を残して, このツールがあなたに大きな助けになることを願っています. ***

ありがとうございます :)

どんなリクエストでも大歓迎です。