package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	//testTimestamp()
	test0()
}

func test() {
	fmt.Println(fmt.Sprintf(`ISODate("%s")`, time.Now().Format(time.RFC3339)))
	fmt.Println(fmt.Sprintf(`ISODate("%s")`, time.Now().AddDate(0, 0, -1).Format(time.RFC3339)))

	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05-0700"))

	println("??", 30*24*60*60)
}

func test0() {
	t := time.Now()
	ts1 := t.UnixNano() / 1000000
	ts2 := t.Unix()*1e3 + int64(t.Nanosecond())/1e6
	fmt.Println(ts1 == ts2)
	fmt.Println(ts1)
	fmt.Println(ts2)
}

func testTimestamp() {
	t := time.Unix(1614305184, 0)
	fmt.Println(t.String())
	// 2021-02-26 10:06:24 +0800 CST
	// 2021-02-26 10:06:24
	a := time.Now().Sub(t)
	fmt.Printf("", a.String())
}
