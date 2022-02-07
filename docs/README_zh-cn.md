
<h3 align="right">
<a href="https://github.com/wrapTMDB/wrapTMDB_go">Github page</a> |
<a href="https://pkg.go.dev/github.com/wrapTMDB/wraptmdb_go">pkg.go</a>  
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

# [wrapTMDB](https://github.com/wrapTMDB/wrapTMDB) 是什么?

```wrapTMDB``` 是一个包装器集合，并且从TMDB的文档中包装  API 以不同的程序语言实现。
它可以帮助开发者向TMDB请求电影或电视节目的信息和元数据。<br/>

这个库由 Golang 编写并在 pkg.go 中发布,<br/>
查看[更多](https://github.com/wrapTMDB/wrapTMDB).
___
## 什么样的项目适合使用？

- 如果您想制作一个客户端能够追踪新电影信息。
- 如果您想制作一个工具来帮助您管理电影文件或视频。
- 甚至你想制作一个程序取代TMDB的官方网站。 ((笑
- ...

___
## 使用方式

### Install:

```bash
$go get github.com/wrapTMDB/wrapTMDB_go
```

在使用此工具之前，请确保您已经拥有 [api_key](https://developers.themoviedb.org/3/getting-started/authentication).
<br/>

``` Golang
package main

import (
	"fmt"
	wraptmdb "wraptmdb_go"
)

func main() {
	//initialize
	wraptmdb.Init("Your api_key")
	wraptmdb.SetHeader(map[string]string{
		"User-Agent": "wraptmdb_go dev",
		"Referer":    "wraptmdb-go",
	})
	//call function
	msg := wraptmdb.Movies.GetDetail("624860", "")

	fmt.Print(msg)
}
```
___

## 我该如何识别这些 API ?

### 运用你的直觉:
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



# 加入开发 ?
```bash
$git clone https://github.com/wrapTMDB/wrapTMDB_go &&

npm install ||

touch src/index.ts 
```

___
## 其他

*** 留个星星吧，希望这个工具能给你很大的帮助。 ***

THANK YOU :)

欢迎任何要求。