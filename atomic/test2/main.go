package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	for i := 0; i < 10; i++ {
		//test1()
		test2()
	}
}

func test1() {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum += 1 //1
		}()
	}
	wg.Wait()
	fmt.Println("test1:", sum)
}

func test2() {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&sum, 1) //2
		}()
	}
	wg.Wait()
	fmt.Println("test2:", sum)
}
