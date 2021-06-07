package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	test()
	test1()
	test2()
}

func test() {
	m := 20 % 10
	fmt.Println(m)
}

func test1() {
	var current int32 = 1
	n := atomic.AddInt32(&current, int32(2147483647))
	m := int(n)
	fmt.Println(m)
}

func test2() {
	var current uint32 = 2147483647 + 2
	n := atomic.AddUint32(&current, uint32(2147483647))
	m := int(n)
	fmt.Println(m)
}
