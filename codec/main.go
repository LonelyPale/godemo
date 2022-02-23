package main

import "encoding/base64"

func main() {
	str := base64.RawStdEncoding.EncodeToString([]byte("abcdef-123456-你好杭州"))
	println(str)
}
