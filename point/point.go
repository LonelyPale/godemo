package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func main() {
	u1 := (*User)(nil)
	u2 := &User{}
	fmt.Println(u1 == u2)
	fmt.Printf("u1=%v\n", u1)
	fmt.Printf("u2=%v\n", u2)
	fmt.Println("===== =====")

	t1 := reflect.TypeOf(u1)
	t2 := reflect.TypeOf(u2)
	fmt.Println(t1 == t2)
	fmt.Println("===== =====")

	v1 := reflect.New(t1.Elem()).Interface()
	v2 := reflect.New(t2.Elem()).Interface()
	fmt.Println(v1 == v2)
	fmt.Println(v1)
	fmt.Println(v2)
}
