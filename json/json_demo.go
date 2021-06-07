package main

import (
	"fmt"
	"time"

	"github.com/json-iterator/go/extra"
	log "github.com/sirupsen/logrus"

	"github.com/LonelyPale/goutils/encoding/json"
)

// -tags jsoniter
func init() {
	extra.SupportPrivateFields()
}

type User struct {
	Name   string
	Data   interface{} `json:"data,omitempty"`
	Time   time.Time   `json:"time"`
	Number int
	str    string
	Sub
	ID int `json:"id,omitempty"`
}

type Sub struct {
	ID int    `json:"id,omitempty"`
	SN string `json:"sn,omitempty"`
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}

func test1() {
	fmt.Println("?nil==nil", &[]interface{}{nil}[0] == nil)

	var t *int
	a := make([]interface{}, 1)
	a[0] = t
	fmt.Println("?-1", a[0], a[0] == nil) //ps.1

	var b interface{}
	var tt *int = nil
	b = tt
	//b = nil
	fmt.Println("?-2", b, tt, b == nil, tt == nil) //ps.2
	fmt.Println("?-2.1", b == tt)                  //ps.2.1
	//fmt.Println("?-2.2", b.(*int) == nil, b.(interface{}) == nil)

	user := User{
		Name: "json",
		//Data: "123",
		Data: &[]interface{}{nil}[0],
		str:  "abc",
	}

	data, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	u := new(User)
	//str := `{"Name":"","Data":[1,2,3]}`
	str := `{"Name":"jj","Data":{"a":1, "B":"asdf", "C":true}}`
	if err := json.Unmarshal([]byte(str), &u); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

func test2() {
	str := `{"number":12.3}`
	var obj User

	if err := json.Unmarshal([]byte(str), &obj); err != nil {
		fmt.Println(err)
	}

	fmt.Println(obj)
}

func test3() {
	str := "123"
	var n int

	if err := json.Unmarshal([]byte(str), &n); err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}

type Message struct {
	Type string
	Data interface{}
}

type Msg struct {
	*Message
	Data json.RawMessage
}

func test4() {
	str := `{"type":"test", "data":123}`
	msg := &Msg{}
	if err := json.Unmarshal([]byte(str), msg); err != nil {
		log.Error(err)
	}
	fmt.Printf("%s %v\n", msg.Type, msg)

	msg.Message.Data = msg.Data

	var data int
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		log.Error(err)
	}
	fmt.Printf("%v\n", data)
}

func test5() {
	user := User{
		Name: "test-user",
		Sub: Sub{
			ID: 123,
			SN: "abc",
		},
	}

	bs, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

func test6() {
	str := `{"Name":"test-user","Number":0,"str":"","id":123,"sn":"abc"}`
	user := &User{}
	if err := json.Unmarshal([]byte(str), user); err != nil {
		panic(err)
	}
	fmt.Println(user)
}
