package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	time.Sleep(3 * time.Second)
	end := time.Now().Unix()
	fmt.Println(start)
	fmt.Println(end)
}
