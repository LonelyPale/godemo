package main

import "fmt"

func main() {
	test()

}

func test() {
	m := make(map[int]string, 2)
	fmt.Println(len(m), m)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	fmt.Println(len(m), m)
}
