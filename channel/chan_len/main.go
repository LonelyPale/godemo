package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go func() {
		c <- 1
		c <- 2
	}()

	time.Sleep(time.Second)
	n := len(c)
	fmt.Println(n)
}
