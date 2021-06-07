package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(1, "start-recover", r)
			panic("test panic")
			fmt.Println(2, "end-recover")
		}
	}()

	panic(123)
}
