package main

import (
	"fmt"
	"hash/crc32"
)

func crc32String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

//生产hash值
func getHashCode(key string) int {
	if len(key) == 0 {
		return 0
	}

	chars := []rune(key)
	lastIndex := len(chars) - 1
	hash := 0
	for i := range chars {
		if i == lastIndex {
			hash += int(chars[i])
		}
		//31 * i ==  (i << 5) - i
		//更好的分散hash
		hash += (hash + int(chars[i])) * 31
	}
	return hash
}

//自定义散列函数(得出存放位置)
func hashCode(k string, length int) int {
	sum := 0
	b := []byte(k)
	for _, v := range b {
		sum += int(v)
	}
	return sum % length
}

func main() {
	str := "123456abcd-1231231111111111111111111111111111111111111111asdfqwexzc13"

	for i := 0; i < 3; i++ {
		hc := crc32String(str)
		fmt.Println("hashcode:", hc)
	}

	for i := 0; i < 3; i++ {
		hc := getHashCode(str)
		fmt.Println("hashcode:", hc)
	}

	for i := 0; i < 3; i++ {
		hc := hashCode(str, 10)
		fmt.Println("hashcode:", hc)
	}
}
