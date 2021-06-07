package main

type User interface {
	String() string
	Man
}

type Man interface {
	Name() string
	Age() int
}

type user struct {
	name string
	age  int
}

func (user) String() {
}

func (user) Name() {
}

func (user) Age() {
}

func main() {
	u := user{
		name: "json",
		age:  10,
	}

	u.Man()
}
