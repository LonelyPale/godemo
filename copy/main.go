package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 3)
	b := []int{1, 2, 3}
	copy(a, b)
	fmt.Print(a, b)
}
