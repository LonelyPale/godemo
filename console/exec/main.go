// go语言运行shell命令
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	shell "github.com/kballard/go-shellquote"
)

func main() {
	ping()
}

func test() {
	str := `nohup bash -c 'cd /root/app/chia-blockchain && . ./activate && sleep  0m && chia plots create -n 1 -k 32 -b 6000 -r 4 -t /tmp1/1 -2 /tmp1 -d /plot3 -f 80e5f08c28219161354672887aab11e8c7c5774fd2670c8b48b61abdf2739d8dbaaf7aa91bcba7e0b9f5ce09f26445a7 -p 97ac546d80383cede84e85fc504c012753ed223f72408eff7e94683085078ed4560eb917bb458f72a03ad1e6773ac8d9' > $LOG_PATH/plot1.log 2>&1 &`

	arr := strings.Fields(str)
	fmt.Println(arr)

	splitArgs, err := shell.Split(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(splitArgs)
}

func ping() {
	var ip, whoami []byte
	var err error
	var cmd *exec.Cmd

	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("ping", "-c 3", "localhost")
	if whoami, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}
	// 默认输出有一个换行
	fmt.Println(string(whoami))
	// 指定参数后过滤换行符
	fmt.Println(strings.Trim(string(whoami), "\n"))

	fmt.Println("====")

	// mac平台获取ip地址
	// 执行连续的shell命令时, 需要注意指定执行路径和参数, 否则运行出错
	cmd = exec.Command("/bin/sh", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(ip))
	fmt.Println(strings.Trim(string(ip), "\n"))
}

func ping2() {
	cmd := exec.Command("ping", "-c 3", "localhost")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
}
