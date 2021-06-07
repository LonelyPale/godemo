package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//cmd := exec.Command("go", "run", "$GOROOT/src/crypto/tls/generate_cert.go", `--host="localhost"`)
	if pwd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		fmt.Println(pwd)
	}

	//cmd := exec.Command("go", "version >1.txt")
	cmd := exec.Command("/bin/bash", "-c", `df -lh`)

	// 执行命令，返回命令是否执行成功
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}

	// 执行命令，并返回结果
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
