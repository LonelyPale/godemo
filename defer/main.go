package main

import "fmt"

func main() {
	test()
	//test1()
	//test2()
}

func test() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func test1() {
	var nums = [5]int{1, 2, 3, 4, 5}
	for i := range nums {
		defer func() { fmt.Println(i) }()
	}
}

func test2() {
	var nums = [5]int{1, 2, 3, 4, 5}
	for i := range nums {
		n := i
		defer func() { fmt.Println(n) }()
	}
}
