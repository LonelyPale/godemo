package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

type NumberDerived interface {
	~int64 | ~float64
}

type ID int64

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	ids := map[string]ID{
		"first":  ID(34),
		"second": ID(12),
	}

	fmt.Printf("非泛型求和: %v and %v\n", SumInts(ints), SumFloats(floats))

	fmt.Printf("泛型求和: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))

	fmt.Printf("泛型求和, 参数类型推断: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	fmt.Printf("泛型求和, 类型约束接口: %v and %v\n", SumNumbers(ints), SumNumbers(floats))

	fmt.Printf("泛型求和, 衍生类型约束接口: %v\n", SumNumbersDerived(ids))

	ForEach([]string{"你好，", "泛型！"}, func(s string) {
		fmt.Printf(s)
	})

	fmt.Println()

	ForEachWithInterface([]any{"你好，", "泛型！"}, func(s any) {
		fmt.Printf(s.(string))
	})
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbersDerived[K comparable, V NumberDerived](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func ForEach[T any](list []T, action func(T)) {
	for _, item := range list {
		action(item)
	}
}

func ForEachWithInterface(list []any, action func(any)) {
	for _, item := range list {
		action(item)
	}
}
