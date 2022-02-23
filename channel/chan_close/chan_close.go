package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test2()
}

func test() {
	c := make(chan int, 2)
	c <- 1

	i, ok := <-c
	if ok {
		println(i, ok)
	} else {
		println(i, ok, "channel closed")
	}

	c <- 2
	c <- 3

	close(c)
	//close(c)

	for v := range c {
		fmt.Println(v)
	}

	i, ok = <-c
	if ok {
		println(i, ok)
	} else {
		println(i, ok, "channel closed")
	}

	//c <- 3
	fmt.Println(-1, <-c)
	fmt.Println(-2, <-c)
	fmt.Println("end")
}

func test1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(123, r)
		}
	}()

	c := make(chan struct{})
	close(c)
	close(c)

	fmt.Println("end")
}

type User struct {
	Name     string
	userChan chan *User
	quit     chan struct{}
	quitMu   sync.RWMutex
}

func (u *User) close() {
	u.quitMu.Lock()
	defer u.quitMu.Unlock()

	select {
	case <-u.quit:
	default:
		close(u.quit)
	}

	fmt.Println("closed")
}
func (u *User) isClose() bool {
	u.quitMu.RLock()
	defer u.quitMu.RUnlock()

	select {
	case <-u.quit:
		return true
	default:
		return false
	}
}

func (u *User) print(i int) {
	//u.quitMu.RLock()
	//defer u.quitMu.RUnlock()

	fmt.Println("running", i, "start")

	select {
	case <-u.quit:
		fmt.Println("closed", i)
		return
	default:
		time.Sleep(time.Second)
	}

	fmt.Println("running", i, "end")
}

func test2() {
	user := &User{quit: make(chan struct{}), userChan: make(chan *User, 1)}
	user.print(1)

	go user.print(2)

	go func() {
		time.Sleep(time.Millisecond * 300)
		user.close()
	}()

	time.Sleep(time.Second * 3)

	user.userChan <- &User{Name: "abc"}
	close(user.userChan)
	fmt.Printf("%v\n", <-user.userChan)
	fmt.Printf("%v\n", <-user.userChan)
}

//结果
//running 1 start
//running 1 end
//running 2 start
//closed
//running 2 end
