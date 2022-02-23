package main

import (
	"fmt"
	"runtime"
)

// go run -gcflags=-G=3 main.go
func main() {
	fmt.Println(runtime.Version())
	printAny(123)
	printAny("abc")
}

func printAny[T any](t T) {
	fmt.Println(t)
}
