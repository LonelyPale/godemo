package main

import (
	"fmt"
	"github.com/LonelyPale/goutils/errors"
)

func main() {
	err := errors.New("error")
	err1 := errors.Wrap(err, "Wrap err1")
	err11 := errors.Wrap(err1, "Wrap err11")
	fmt.Println(err1)
	fmt.Println(err11)

	err2 := errors.WithMessage(err, "WithMessage err2")
	err21 := errors.WithMessage(err2, "WithMessage err21")
	fmt.Println(err2)
	fmt.Println(err21)
}
