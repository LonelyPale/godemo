package main

import (
	"fmt"
	"reflect"
)

type empty struct{}

func main() {
	printKind(empty{})

	e := &empty{}
	printKind(&e)
}

func printKind(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr:
		fmt.Println("Ptr")
	case reflect.Struct:
		fmt.Println("Struct")
	default:
		fmt.Println("??")
	}
}
