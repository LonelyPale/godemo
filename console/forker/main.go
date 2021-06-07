package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//sleep()
	//ping()
	bash()
	//nohup()
	//forker()
}

func sleep() {
	cmd := exec.Command("sleep", "10")
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("sleep:", cmd.Process.Pid)
	//time.Sleep(time.Second)
	//os.Exit(1)
}

func ping() {
	cmd := exec.Command("ping", "-c88", "localhost")
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("ping:", cmd.Process.Pid)
}

func bash() {
	path := "/Users/wyb/project/github/godemo/console/forker/tmp2.log"
	//cmd := exec.Command("bash", "-c", "ping -c88 localhost")//能启动，无输出，pid正常
	//cmd := exec.Command("bash", "-c", "ping -c88 localhost", ">", path)//能启动，无输出，pid正常
	cmd := exec.Command("bash", "-c", "ping -c88 localhost > "+path) //能启动，有输出，但pid变了
	//cmd := exec.Command("bash", "-c", "ping -c88 localhost | tee "+path)//能启动，有输出，但pid变了
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("bash:", cmd.Process.Pid)
}

func nohup() {
	//path := "/Users/wyb/project/github/godemo/console/forker/tmp3.log"
	//cmd := exec.Command("nohup", "ping", "-c88", "localhost", ">", "/dev/null", "2>&1", "&")//有bug，不能启动
	cmd := exec.Command("nohup", "ping", "-c88", "localhost")
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("nohup:", cmd.Process.Pid)
}

func forker() {
	path := "/Users/wyb/project/github/godemo/console/forker/tmp4.log"
	//cmd := exec.Command("nohup", "bash", "-c", "ping -c88 localhost", ">", "/dev/null", "2>&1", "&")
	//cmd := exec.Command("nohup", "bash", "-c", "ping -c88 localhost", ">", path, "2>&1", "&") //有bug，能启动，无输出
	cmd := exec.Command("nohup", "bash", "-c", "ping -c88 localhost > "+path) //能启动，但pid变了
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("forker:", cmd.Process.Pid)
}
