package main

import "fmt"

func main() {
	test3()
}

func test() {
	arr := []string{"a", "b", "c", "d", "e", "f"}

	for i, v := range arr {
		fmt.Println(i, v)
		if v == "c" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	fmt.Println(arr)
}

func test1() {
	arr := []string{"a", "b", "c", "d", "e", "f"}

	for i := 0; i < len(arr); {
		fmt.Println(i, arr[i])
		if arr[i] == "c" {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}

	fmt.Println(arr)
}

func test2() {
	arr := []string{"a", "b", "c", "d", "e", "f"}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(i, arr[i])
		if arr[i] == "c" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	fmt.Println(arr)
}

func test3() {
	arr := []int{1, 2, 3, 4, 5}
	printArr(arr)
}

func printArr(arr []int) {
	fmt.Printf("len=%v, cap=%v, %v", len(arr), cap(arr), arr)
}
