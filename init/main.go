package main

import "fmt"

type magicSecretKey struct {
	a string
	b int
}

var Mage magicSecretKey

func init() {
	Mage.a = "abc"
	Mage.b = 123
	fmt.Println(Mage)
}

func main() {

}
