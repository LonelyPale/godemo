package main

import "fmt"

func main() {
	//n1 := []int{1, 3, 5, 7, 9}
	//n2 := []int{2, 4, 6, 8, 10}
	expr := 88
	switch expr {
	case 1, 3, 5, 7, 9:
		fmt.Println("odd number")
	case 2, 4, 6, 8, 10:
		fmt.Println("even numbers")
	default:
		fmt.Println("unknown")
	}
}
