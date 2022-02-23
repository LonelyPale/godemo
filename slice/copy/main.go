package main

import (
	"fmt"
	"unsafe"
)

func main() {
	test1()
}

func test1() {
	data := make([]int, 0, 3)

	// 24  len:8, cap:8, array:8
	fmt.Println(unsafe.Sizeof(data))

	// 我们通过指针的方式，拿到数组内部结构的字段值
	ptr := unsafe.Pointer(&data)
	opt := (*[3]int)(ptr)

	// addr, 0, 3
	fmt.Println(opt[0], opt[1], opt[2])

	data = append(data, 123)

	fmt.Println(unsafe.Sizeof(data))

	shallowCopy := data[:1]

	ptr1 := unsafe.Pointer(&shallowCopy)

	opt1 := (*[3]int)(ptr1)

	fmt.Println(opt1[0])
}
