package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var str = "hello 你好"
	fmt.Println(len("你好啊"))
	fmt.Println(len([]rune("你好啊")))

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))

	tmp := strings.Split("888", "/")
	fmt.Println(len(tmp), tmp)
}
