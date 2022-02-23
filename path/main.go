package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	test(".")
	test("..")
}

func test(p string) {
	str, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p, str)
}
