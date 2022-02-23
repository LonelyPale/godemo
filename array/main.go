package main

import "fmt"

type ID [12]byte

func main() {
	var id ID
	fmt.Println(len(id), id)
}
