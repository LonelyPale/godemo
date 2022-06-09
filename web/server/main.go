package main

import "github.com/lonelypale/goutils/net/http"

func main() {
	server := http.NewServer(http.ServerOptions{})
	server.AddRouter(http.NewRouter(nil, nil))
	server.Run()
}
