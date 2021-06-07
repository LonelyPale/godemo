package main

import "fmt"

func main() {
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
