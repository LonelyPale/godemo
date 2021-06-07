package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	info()
	file()
	test()
}

func info() {
	println(`系统类型：`, runtime.GOOS) //sysType: darwin linux windows
	println(`系统架构：`, runtime.GOARCH)
	println(`CPU 核数：`, runtime.GOMAXPROCS(0))

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	println(`电脑名称：`, name)
}

func file() {
	fmt.Println(os.FileMode(777), 777)
	fmt.Println(os.FileMode(0777), 0777)
	fmt.Println(os.FileMode(0755), 0755)
}

func test() {
	device := "/dev/nvme0n1"
	if strings.HasPrefix(device, "/dev/nvme") { // /dev/nvme0n1p1
		fmt.Print(fmt.Sprintf("%sp1\n", device))
	}
}
