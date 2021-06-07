package main

import (
	"fmt"
	"time"
)

const DefaultTimeFormart = "2006-01-02 15:04:05"

func main() {
	test1()
}

func test1() {
	str := "2021-02-14 01:23:32"
	t1, err := time.Parse(DefaultTimeFormart, str)
	if err != nil {
		panic(err)
	}
	t2, err := time.ParseInLocation(DefaultTimeFormart, str, time.Local)
	if err != nil {
		panic(err)
	}
	t3 := t2.UTC()
	fmt.Println(str)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
}
