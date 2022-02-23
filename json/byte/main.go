package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Data []byte
}

func main() {
	user := &User{Data: []byte{86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}}
	bs, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	user2 := new(User)
	if err := json.Unmarshal(bs, user2); err != nil {
		panic(err)
	}
	fmt.Println(user2)
}
