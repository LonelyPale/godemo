package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	test2()
}

func test() {
	out := bytes.NewBufferString("abcdefgh\n1\n2\n3\n")
	buf := make([]byte, 8)

	fmt.Println("out-1:", out.Len(), out.String())
	fmt.Println("buf-1:", len(buf), string(buf))

	//_, err := out.Read(buf)
	t := out
	s, err := out.ReadString('\n')
	fmt.Println(111, t.Len(), s)
	s, err = t.ReadString('\n')
	fmt.Println(222, t.Len(), s)
	s, err = t.ReadString('\n')
	fmt.Println(333, t.Len(), s)

	s, err = t.ReadString('\n')
	fmt.Println(333, t.Len(), s)
	s, err = t.ReadString('\n')
	fmt.Println(333, t.Len(), s)
	s, err = t.ReadString('\n')
	fmt.Println(333, t.Len(), s)
	s, err = t.ReadString('\n')
	fmt.Println(333, t.Len(), s)

	buf = []byte(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("out-2:", out.Len(), out.String())
	fmt.Println("buf-2:", len(buf), string(buf), "=>", s)

	_, err = out.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}

func test1() {
	out := bytes.NewBufferString("")
	buf := make([]byte, 8)
	n, err := out.Read(buf)
	if err != nil {
		fmt.Println(n, err)
	}
	fmt.Println(n, string(buf))
}

func test2() {
	w := bytes.NewBufferString("qwe")
	r := bytes.NewBufferString("123")
	t := bytes.NewBufferString("!@#")

	fmt.Println(1, w)
	fmt.Println(2, r, t)

	if _, err := io.Copy(w, r); err != nil {
		fmt.Println(err)
	}
	fmt.Println(3, w)
	fmt.Println(4, r)

	if _, err := io.Copy(w, t); err != nil {
		fmt.Println(err)
	}
	fmt.Println(5, w)
	fmt.Println(6, t)

	if _, err := io.Copy(w, t); err != nil {
		fmt.Println(err)
	}
	fmt.Println(7, w)
	fmt.Println(8, t)
}
