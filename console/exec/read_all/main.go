package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//test1()
	test2()
}

func test1() {
	cmd := exec.Command("blkid", "/dev/sdf1")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

func test2() {
	cmd := exec.Command("bash", "-c", "echo 'abc 123' >> /Users/wyb/project/github/godemo/console/exec/read_all/tmp.txt")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
