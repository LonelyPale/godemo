package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test()
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
	println("test WaitGroup start...")
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		n := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			println("test WaitGroup:", n)
		}()
	}

	wg.Wait()
	println("test WaitGroup end...")
}
