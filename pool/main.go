package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() (err error, num int) {
	defer func() {
		err = errors.New("3")
	}()

	err = errors.New("1")
	num = 11
	return errors.New("2"), 22
}
