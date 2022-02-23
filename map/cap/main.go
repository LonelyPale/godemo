package main

import "fmt"

func main() {
	test()

}

func test() {
	m := make(map[int]string, 2)
	fmt.Printf("len=%d, %v\n", len(m), m)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	fmt.Printf("len=%d, %v\n", len(m), m)
	delete(m, 1)
	fmt.Printf("len=%d, %v\n", len(m), m)

}
