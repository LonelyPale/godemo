package main

import (
	"fmt"
	"github.com/ThomasRooney/gexpect"
	"strings"
)

func main() {
	ping()
}

func ping() {
	child, err := gexpect.Spawn("ping -c 3 127.0.0.1")
	if err != nil {
		panic(err)
	}
	//child.Interact()

	//ReadAll(child)
	//ReadLine(child)

	ss, s, err := child.ExpectRegexFindWithOutput(".*PING.*")
	fmt.Println("ss:", ss)
	fmt.Println("s:", s)
	if err != nil && !strings.Contains(err.Error(), "ExpectRegex didn't find regex") {
		panic(err)
	}

	ss, err = child.ExpectRegexFind(".*64.*")
	fmt.Println("ss:", ss)
	if err != nil && !strings.Contains(err.Error(), "ExpectRegex didn't find regex") {
		panic(err)
	}

	fmt.Println("Success")
}

func ReadAll(proc *gexpect.ExpectSubprocess) {
	for {
		str, err := proc.ReadUntil(byte(0))
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println(666, string(str))
				fmt.Println("<EOF>")
				break
			} else {
				panic(err)
			}
		}
		fmt.Println(888, str)
	}
}

func ReadLine(proc *gexpect.ExpectSubprocess) {
	for {
		str, err := proc.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println(666, str)
				fmt.Println("<EOF>")
				break
			} else {
				panic(err)
			}
		}
		fmt.Println(888, str)
	}
}
