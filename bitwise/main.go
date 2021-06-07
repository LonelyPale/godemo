package main

import (
	"fmt"
)

func main() {
	//a := 55205414831165500
	a := 631712498580700781
	n := 4
	fmt.Println(a >> n)
	fmt.Println(a << n)
	fmt.Println(exponent(2, 5), exponent(2, 8))
}

//a的n次方
//超出uint64的部分会丢失
func exponent(a, n uint64) uint64 {
	result := uint64(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}
