package main

import (
	"fmt"
	"sync"
)

func main() {
	test2()
}

func test() {
	var wg sync.WaitGroup
	var m = map[int]int{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			m[i] = i
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	fmt.Println("test(10):", len(m))
}

func test1() {
	var m = map[int]int{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1) // 必须放在协程外，不能放在协程内。
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			m[i] = i
		}(i)
	}

	wg.Wait()
	fmt.Println("test1(10):", len(m))
}

func test2() {
	var m = sync.Map{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait()
	var count int
	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Println("test2(10):", count)
}
