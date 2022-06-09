package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"os"
)

const text = `软件简介
安全、稳定、高性能的内网穿透工具，用 Rust 语言编写

rathole，类似于 frp 和 ngrok，可以让 NAT 后的设备上的服务通过具有公网 IP 的服务器暴露在公网上。

高性能 具有更高的吞吐量，高并发下更稳定。见Benchmark
低资源消耗 内存占用远低于同类工具。见Benchmark。二进制文件最小可以到 ~500KiB，可以部署在嵌入式设备如路由器上。
安全性 每个服务单独强制鉴权。Server 和 Client 负责各自的配置。使用 Noise Protocol 可以简单地配置传输加密，而不需要自签证书。同时也支持 TLS。
热重载 支持配置文件热重载，动态修改端口转发服务。HTTP API 正在开发中。
`

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("/tmp/example.bleve", mapping)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll("/tmp/example.bleve")

	// index some data
	err = index.Index("软件简介", text)

	// search for some text
	query := bleve.NewMatchQuery("性能")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		panic(err)
	}

	fmt.Println(searchResults)
	fmt.Println(searchResults.Request.From)
	fmt.Println(searchResults.Hits.Len())
	fmt.Println(searchResults.Status)
	fmt.Println(searchResults.Total)
	fmt.Println(searchResults.Size())
	fmt.Println(searchResults.String())
}
