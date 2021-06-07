package main

import (
	"fmt"
	"reflect"
)

func main() {
	test1()
}

func test1() {
	m := map[string]interface{}{"_id": "abc", "num": 123}

	vDoc := reflect.ValueOf(m)
	var vid reflect.Value
	switch vDoc.Kind() {
	case reflect.Map:
		vid = vDoc.MapIndex(reflect.ValueOf("_id"))
	default:
		panic("error type")
	}

	if vid.IsValid() {
		fmt.Println(vid.Interface())
	}

	fmt.Println("IsValid:", vid.IsValid())
	fmt.Println("IsZero:", vid.IsZero())
	fmt.Println("IsNil:", vid.IsNil())
}
