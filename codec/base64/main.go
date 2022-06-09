package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var str = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAACM0lEQVRYhe1WPUgDMRj9TsRNEKqI\n4lJoHVwEQYviUifB0S4VQRAVnNxcxEVxKTi4iTiKLrrpqktRRBxcRChyDtJS5VzcFIm8aGIuvbPJ\ntSBCH4S7XH7e+17yJeewcp7RH6LpL8mBhoCGgGbjnq+HRK0Z+b6V2fI1Lx0u/bRbwM4BiNBIOXEN\nMHfgG8/uLnW0t31Vps4r2gQ64nNG80U7iDQnJCIsgZUAp3PUqB8r540FGO8BU3Lbvv8jDfWIijvZ\nwH7dCwe+MSZLYZUFYcR6uyqkbgIwuTqxLkZvMxUR2QEQqBFXcydUQCw1zV+8yz2jAWGRie+2QngW\n3BydkRCiA5tJnZQx5nuGfcMYk3TkBxHIIaJ/Mi2dEIJeHh4i2wtXqmUCdwCkIBdOoJTW38k7iVcl\nqBVyE0JELJWmofFhXu9avSAil0egL0O9opcOBMFzXWOiWiAdELYjcpC/7fWQ5z5SLB6vcEHf8UF1\n0wvJdw60DPRy20EKcvpIczGoqxHry1HLeSCzABsuNuH6soAvw+0YUd8pOY7jIwM2clf8ubI86BNI\nFley/B/ghMphFJTD4oj97TIKO4ZDBUGAWvAJ5elun13MbvKSSCSYQHEny9vFU33HUwBjxHjMJfrq\nfL49gKivF0cqRJ4fr1EymaRCoUBd8/tU1Ha/iBxtAPpizH2uJPuIefVr2uiHZGZ2W4oAuAhlGYLI\nMcYE0oGw6INEqE4IQabk4FBdsLqOg0QI2EbOQUSfLc14jhB7b40AAAAASUVORK5CYII="

func main() {
	res, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("test.jpg", res, 0777); err != nil {
		panic(err)
	}

	hash1 := md5.Sum(res)
	fmt.Printf("%x\n", hash1)

	png, err := ioutil.ReadFile("/Users/wyb/temp/001/稀有英雄100个/像素/CryptoSanguo0001.png")
	if err != nil {
		panic(err)
	}

	hash2 := md5.Sum(png)
	fmt.Printf("%x\n", hash2)

	if hash1 != hash2 {
		panic("!=")
	} else {
		fmt.Println("==")
	}
}
