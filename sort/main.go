package main

import (
	"fmt"
	"sort"
)

func main() {
	test1()
	test2()
}

func test1() {
	arr := []int{1, 3, 5, 7}
	//arr := []int{1, 2, 3, 4, 5}
	//arr := []int{3, 4, 1, 2, 5, 6, 1, 3, 1, 0, 3, 5, 7, 0, -1}
	//arr := []int{-1, 0, 0, 1, 1, 1, 2, 3, 3, 3, 4, 5, 5, 6, 7}
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] == 3
	})

	fmt.Println(index)
}

func test2() {
	ts := 400 - 200
	arr := []int{100, 150, 200, 250, 300}
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] > ts
	})

	fmt.Println(index)
}
