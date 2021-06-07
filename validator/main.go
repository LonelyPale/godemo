package main

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	Num        int       `validate:"min=10"`
	Number     string    `validate:"omitempty,latitude"` // number numeric latitude longitude
	CreateTime time.Time `validate:"omitempty,datetime" vCreate:"isdefault" vModify:"isdefault" label:"创建时间"`
	ModifyTime time.Time `validate:"omitempty,datetime" vCreate:"isdefault" vModify:"isdefault" label:"修改时间"`
}

var timeFun validator.Func = func(fl validator.FieldLevel) bool {
	_, ok := fl.Field().Interface().(time.Time)
	return ok
}

func test1() {
	validate := validator.New()
	validate.SetTagName("vCreate")
	err := validate.Struct(User{})
	if err != nil {
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

	err := validate.Struct(User{88, "443.3", time.Now(), time.Now()})
	if err != nil {
		panic(err)
	}
}

func main() {
	test1()
	test2()
}
