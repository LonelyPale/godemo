package main

import (
	"fmt"
	"reflect"
)

func main() {
	nilType := reflect.TypeOf(nil)
	fmt.Println(nilType)
	fmt.Println(nilType.Kind())

	nilValue := reflect.ValueOf(nil)
	val := reflect.Indirect(nilValue)
	fmt.Println(val)
}
