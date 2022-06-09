package main

import (
	"fmt"
	"strconv"
)

func main() {
	println(strconv.FormatFloat(77, 'e', 2, 64))
	println(strconv.FormatFloat(77, 'f', -1, 64))

	var f float64
	f = 1234.5678
	fmt.Println(strconv.FormatFloat(f, 'f', 2, 64))  // "1234.56"
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 64)) // "1234.5678"
}
