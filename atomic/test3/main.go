package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	test()
}

func test() {
	lb := RoundRobinBalancer{
		backends: []bool{true, false, false, true, true},
		current:  4,
	}

	m := NewMap()
	var wg sync.WaitGroup
	for i := 0; i < 12000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			idx := lb.NextPeer()
			//fmt.Print(idx, " ")
			m.Add(idx)
		}(i)
	}
	wg.Wait()

	fmt.Println()
	m.Print()
}

type RoundRobinBalancer struct {
	backends []bool
	current  uint64
	mux      sync.Mutex
}

func (r *RoundRobinBalancer) NextIndex() int {
	return int(atomic.AddUint64(&r.current, uint64(1)) % uint64(len(r.backends)))
}

// NextPeer 返回下一个可用的服务器
func (r *RoundRobinBalancer) NextPeer() int {
	//r.mux.Lock()
	//defer r.mux.Unlock()

	// 遍历后端列表，找到可用的服务器
	next := r.NextIndex()
	l := len(r.backends) + next // 从 next 开始遍历
	for i := next; i < l; i++ {
		idx := i % len(r.backends) // 通过取模计算获得索引
		// 如果找到一个可用的服务器，将它作为当前服务器。如果不是初始的那个，就把它保存下来
		if r.backends[idx] {
			if i != next {
				atomic.StoreUint64(&r.current, uint64(idx)) // 标记当前可用服务器
				//r.current = uint64(idx)
			}
			return idx
		}
	}
	return -1
}

type Map struct {
	m   map[int]int
	mux sync.Mutex
}

func NewMap() *Map {
	return &Map{
		m: make(map[int]int),
	}
}

func (m *Map) Add(n int) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.m[n] = m.m[n] + 1
	//if a, ok := m.m[n]; ok {
	//	m.m[n] = a + 1
	//} else {
	//	m.m[n] = 1
	//}
}

func (m *Map) Print() {
	for key, value := range m.m {
		fmt.Println("test:", key, value)
	}
}
