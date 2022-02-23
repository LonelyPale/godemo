package main

import (
	"fmt"
	"unsafe"
)

func main() {
	test1()
}

type Book struct {
	Title string
}

type User struct {
	Name string
	Age  int
	*Book
}

func test() {
	p1 := &Book{}
	p2 := User{
		Name: "abc",
		Age:  123,
	}

	fmt.Println(unsafe.Sizeof(&p1))
	fmt.Println(unsafe.Sizeof(&p2))
}

func test1() {
	user := &User{Name: "abc", Age: 123}
	printUser(*user)
	printUser(*user)
	fmt.Println(*user == *user)
}

func printUser(user User) {
	fmt.Printf("%p, %v\n", &user, user)
}
