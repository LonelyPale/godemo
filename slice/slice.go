package main

import (
	"bytes"
	"fmt"
)

var (
	slice = []string{"", "", "a", "", "b", "", "c"}
)

func test1() {
	elems := make([]string, len(slice))
	copy(elems, slice)
	for i, a := range elems {
		if len(a) == 0 {
			elems = append(elems[:i], elems[i+1:]...)
		}
	}

	fmt.Println(len(elems), elems)
}

func test2() {
	elems := make([]string, 0)
	for _, s := range slice {
		if len(s) > 0 {
			elems = append(elems, s)
		}
	}

	fmt.Println(len(elems), elems)
}

func test3() {
	failures := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("test3-init:", failures)
	for i, v := range failures {
		failures = failures[1:]
		fmt.Println("test3:", i, v, failures)
	}
}

func test4() {
	slice := make([]string, 3)
	fmt.Println("len =", len(slice))

	slice[0] = "a"
	slice[2] = "c"

	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func test5() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	b[0] = 10
	c := a
	fmt.Printf("c:%p %v\n", c, c)
	c = append(c, 0)
	testSlice(b)
	fmt.Printf("a:%p %v\n", a, a)
	fmt.Printf("b:%p %v\n", b, b)
	fmt.Printf("c:%p %v\n", c, c)
}

func testSlice(s []int) {
	s = append(s, 7, 8, 9)
	fmt.Printf("b:%p %v\n", s, s)
}

func test6() {
	a := []byte{1, 2, 3}
	b := []byte{1, 2, 3}
	c := []byte{1, 2, 3, 0}
	fmt.Println("a==b", bytes.Equal(a, b))
	fmt.Println("a==c", bytes.Equal(a, c))
}

//重点注意
func test7() {
	a := []int{1, 2, 3, 4, 5, 6}
	a1 := a[:3]
	a2 := a[3:]
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a1, len(a1), cap(a1))
	fmt.Println(a2, len(a2), cap(a2))
	a1 = append(a1, 8)
	//a2 = append(a2, 9)
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	a1 = append(a1, 0)
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
}

func test8() {
	var a []int
	fmt.Println(a, a == nil)
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	fmt.Println(append([]int{}, 1, 2, 3))
	fmt.Println([]int{} == nil, []byte("") == nil)
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	test8()
}
