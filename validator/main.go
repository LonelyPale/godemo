package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	Num        int       `validate:"min=10"`
	Number     string    `validate:"omitempty,latitude"` // number numeric latitude longitude
	CreateTime time.Time `validate:"omitempty,datetime" vCreate:"isdefault" vModify:"isdefault" label:"创建时间"`
	ModifyTime time.Time `validate:"omitempty,datetime" vCreate:"isdefault" vModify:"isdefault" label:"修改时间"`
	Address    string    `validate:"omitempty" vCreate:"required"`
	Nickname   string    `validate:"omitempty,alphanumunicode,max=30"`
}

var timeFun validator.Func = func(fl validator.FieldLevel) bool {
	_, ok := fl.Field().Interface().(time.Time)
	return ok
}

func test1() {
	validate := validator.New()

	user := User{Num: 123}
	//if err := validate.Struct(user); err != nil {
	//	panic(err)
	//}

	validate.SetTagName("vCreate")
	if err := validate.Struct(user); err != nil {
		panic(err)
	}
}

func test2() {
	validate := validator.New()
	validate.SetTagName("vModify")
	validate.SetTagName("validate")
	if err := validate.RegisterValidation("datetime", timeFun); err != nil {
		panic(err)
	}

	err := validate.Struct(User{88, "90", time.Now(), time.Now(), "地址", "测试昵称abc123①②③"})
	if err != nil {
		panic(err)
	}
}

// MyStruct ..
type MyStruct struct {
	String string `validate:"is-awesome"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func test3() {
	validate = validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyVal)

	s := MyStruct{String: "awesome"}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.String = "not awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}

func main() {
	//test1()
	//test2()
	test3()
}
