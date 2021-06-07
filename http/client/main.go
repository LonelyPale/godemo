package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 3 * time.Second}
	//resp, err := client.Get("http://127.0.0.1:8080/ping")
	resp, err := client.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("response:", resp.StatusCode, resp.Status)
}
