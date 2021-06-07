package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func main() {
	//pipe()
	//buff()
	//ls()
	//ps()
	sudo()
}

func buff() {
	input := "foo  bar   baz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func pipe() {
	reader, writer := io.Pipe()
	inputData := []byte("1234567890ABCD")
	go writer.Write(inputData)
	outputData := make([]byte, 11)
	n, _ := reader.Read(outputData)
	fmt.Println(string(outputData))
	fmt.Println("read number", n)
	fmt.Println(string(outputData))
}

func sudo() {
	if err := execCmd("sudo", []string{"-Sk", "ls", "-lla", "/Users/wyb/project/github/godemo/"}); err != nil {
		fmt.Println(err)
	}
	if err := execCmd("bash", []string{"-c", "sudo -Sk ls -lla /Users/wyb/project/github/godemo/"}); err != nil {
		fmt.Println(err)
	}
}

func ls() {
	if err := execCmd("ls", []string{"-la", "/Users/wyb/project/github/godemo/console"}); err != nil {
		fmt.Println(err)
	}
}

func ps() {
	if err := execCmd2("ps", []string{"aux"}); err != nil {
		fmt.Println(err)
	}
}

func execCmd(shell string, raw []string) error {
	cmd := exec.Command(shell, raw...)

	//stdin, err := cmd.StdinPipe()
	//if err != nil {
	//	return err
	//}
	stdin := &bytes.Buffer{}
	cmd.Stdin = stdin

	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	return err
	//}
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	//stderr, err := cmd.StderrPipe()
	//if err != nil {
	//	return err
	//}

	//#1
	go func() {
		if _, err := io.WriteString(stdin, "wyb123456\n"); err != nil {
			fmt.Println(err)
		}
		fmt.Println("stdin ok.")
	}()

	//#2
	//go func() {
	//	inputBuf := bufio.NewWriter(stdin)
	//	defer inputBuf.Flush()
	//	if _, err := inputBuf.WriteString("wyb123456\n"); err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	go func() {
		//使用带缓冲的读取器
		outputBuf := bufio.NewReader(stdout)
		for {
			output := make([]byte, 1024)
			n, err := outputBuf.Read(output)
			if err != nil {
				// 判断是否到文件的结尾了否则出错
				if err.Error() == "EOF" { //
					if cmd.ProcessState != nil && cmd.ProcessState.Exited() {
						fmt.Printf("2:%s\n", string(output))
						return
					}
				} else {
					fmt.Printf("Error :%s\n", err)
					return
				}
			}
			if n > 0 {
				fmt.Printf("1:%s\n", string(output))
			}
		}
	}()

	//go func() {
	//	//使用带缓冲的读取器
	//	outputBuf := bufio.NewReader(stderr)
	//	for {
	//		//一次获取一行,_ 获取当前行是否被读完
	//		//output, _, err := outputBuf.ReadLine()
	//		output := make([]byte, 1024)
	//		_, err := outputBuf.Read(output)
	//		if err != nil {
	//			// 判断是否到文件的结尾了否则出错
	//			if err.Error() == "EOF" {
	//				if cmd.ProcessState != nil && cmd.ProcessState.Exited(){
	//					fmt.Printf("err2: %s\n", string(output))
	//					return
	//				}
	//			} else {
	//				fmt.Printf("err3: Error: %s\n", err)
	//				return
	//			}
	//		}
	//		fmt.Printf("err1: %s\n", string(output))
	//	}
	//}()

	if err := cmd.Start(); err != nil {
		return err
	}

	//time.Sleep(time.Second)
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 1)

	return nil
}

func execCmd2(shell string, raw []string) error {
	cmd := exec.Command(shell, raw...)

	//stdin, err := cmd.StdinPipe()
	//if err != nil {
	//	return err
	//}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		text := s.Text()
		fmt.Println(text)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}

	return nil
}
