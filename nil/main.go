package main

import "fmt"

type Filter func()

func main() {
	var f Filter
	fmt.Println(f == nil, f)
}
