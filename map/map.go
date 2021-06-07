package main

import "fmt"

type Map map[string]interface{}

func main() {
	//test2()
	test3()
}

func test1() {
	m := Map{"name": "bob", "age": 10}
	sex, ok := m["sex"]
	fmt.Println(sex, ok)

	if sex, ok := m["sex"]; !ok || len(sex.(string)) == 0 {
		fmt.Println("Sex does not exist")
	}

	switch val := m["sex"].(type) {
	case nil:
		fmt.Println("sex is nil")
	case string:
		if len(val) == 0 {
			fmt.Println("sex len is 0")
		}
	default:
		fmt.Println("Unknown type")
	}
}

func test2() {
	m := map[int]bool{1: true, 2: true}
	a, ok := m[0]
	fmt.Println(a, ok)
}

func test3() {
	m1 := map[string]string{"a": "1", "b": "2", "c": "3"}
	mm := interface{}(m1)
	m := mm.(map[interface{}]interface{})
	fmt.Println(m)
}
