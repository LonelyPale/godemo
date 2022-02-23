package main

import (
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/traefik/yaegi/stdlib/unrestricted"
)

const src = `
package foo

import (
	"fmt"
)

func Bar(s string) string {
	return s + "-Foo"
}
`

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err := i.Eval(`import "fmt"`)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(`fmt.Println("Hello Yaegi.\n你好，中国！")`)
	if err != nil {
		panic(err)
	}
}

func test2() {
	i := interp.New(interp.Options{
		GoPath: "/Users/wyb/lib/GOPATH",
	})
	i.Use(stdlib.Symbols)
	i.Use(unrestricted.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar := v.Interface().(func(string) string)

	r := bar("Kung")
	println(r)
}

func test3() {
	i := interp.New(interp.Options{GoPath: "/Users/wyb/project/github/godemo/script/go-script/yaegi/_gopath/"})
	if err := i.Use(stdlib.Symbols); err != nil {
		panic(err)
	}

	if _, err := i.Eval(`import "github.com/foo/bar"`); err != nil {
		panic(err)
	}

	val, err := i.Eval(`bar.NewFoo`)
	if err != nil {
		panic(err)
	}

	fmt.Println(val.Call(nil))
}
