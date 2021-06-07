package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test1()
	fmt.Println("sleep start")
	time.Sleep(time.Second)
}

func test1() {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	fmt.Println("fun end")
}

func test() {
	a := sync.WaitGroup{}
	a.Wait()
}
