package main

import "fmt"

// some interface
type Stringer interface {
	String() string
}

// a struct that implements Stringer interface
type Struct1 struct {
	field1 string
}

func (s Struct1) String() string {
	return s.field1
}

// another struct that implements Stringer interface, but has a different set of fields
type Struct2 struct {
	field1 []string
	dummy  bool
}

func (s Struct2) String() string {
	return fmt.Sprintf("%v, %v", s.field1, s.dummy)
}

// container that can embedd any struct which implements Stringer interface
type StringerContainer struct {
	Stringer
}

func main() {
	// the following prints: This is Struct1
	fmt.Println(StringerContainer{Struct1{"This is Struct1"}})
	// the following prints: [This is Struct1], true
	fmt.Println(StringerContainer{Struct2{[]string{"This", "is", "Struct1"}, true}})
	// the following does not compile:
	// cannot use "This is a type that does not implement Stringer" (type string)
	// as type Stringer in field value:
	// string does not implement Stringer (missing String method)
	//fmt.Println(StringerContainer{"This is a type that does not implement Stringer"})
}
