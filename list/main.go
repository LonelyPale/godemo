package main

import (
	"container/list"
	"fmt"
)

func main() {
	test2(nil)
}

func test() {
	fmt.Println(len(make([]string, 0)))
	fmt.Println(len(make([]string, 3)))

	fmt.Println(cap(make([]string, 0)))
	fmt.Println(cap(make([]string, 3)))

	a := make([]string, 0)
	a = append(a, "1")
	fmt.Println(len(a))

	index := 0
	b := make([]string, 3)
	b = append(b[:index], b[index+1:]...)
	fmt.Println(len(b))
}

func test1() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(17)
	l.PushBack(27)
	l.PushBack(71)
	l.PushBack(74)
	//fmt.Println(l.Back().Value)
	//fmt.Println(l.Front().Value)
	for p := l.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}

func test2(s ...[]int) {
	fmt.Println(s)
	fmt.Println(append(make([][]int, 0), s...))

	for _, v := range s {
		fmt.Println(v)
	}
}
