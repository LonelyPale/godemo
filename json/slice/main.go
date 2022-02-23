package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `["abc123"]`
	var ids []string
	if err := json.Unmarshal([]byte(str), &ids); err != nil {
		panic(err)
	}
	fmt.Println(ids)
}
