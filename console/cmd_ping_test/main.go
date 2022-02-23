package main

import "github.com/LonelyPale/goutils/cmd"

// GOOS=linux go build
func main() {
	str := `bash -c 'ping -c 88 localhost >ping.log 2>&1 &'`
	if err := cmd.Exec(str); err != nil {
		panic(err)
	}
}
