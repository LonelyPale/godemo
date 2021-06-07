package main

import (
	"fmt"
	"github.com/ThomasRooney/gexpect"
	"github.com/google/goterm/term"
)

func main() {
	str, err := RunCmd("bash -c 'blkid /dev/sdf1'")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

func RunCmd(cmd string) (string, error) {
	fmt.Println(term.Blue(cmd))

	child, err := gexpect.Spawn(cmd)
	if err != nil {
		return "", err
	}

	out, err := child.ReadUntil(0x00)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}

	return string(out), nil
}
