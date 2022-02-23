package main

import "fmt"

type User struct {
	Man
	ID string
}

type Man struct {
	Name string
	Age  int
}

func main() {
	user := new(User)
	test1(user)
	test2(user)
}

func test1(obj interface{}) {
	man, ok := obj.(Man)
	fmt.Println(ok, man)
}

func test2(obj interface{}) {
	switch i := obj.(type) {
	case Man:
		fmt.Println(1)
	case *Man:
		fmt.Println(2)
	default:
		fmt.Println(i)
	}
}
