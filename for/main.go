package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		for j := 11; j <= 15; j++ {
			fmt.Printf("%d ", j)
			if j == 13 {
				break
			}
		}
		fmt.Println()
	}
}
