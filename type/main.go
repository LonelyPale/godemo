package main

import "fmt"

func main() {
	var val interface{}
	str := "abc"
	val = &str
	switch v := val.(type) {
	case string:
		fmt.Println("string:", v)
	case *string:
		fmt.Println("*string:", *v)
	default:
		fmt.Println("unknown type")
	}
}
