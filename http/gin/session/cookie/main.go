package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Set("test", `首先调用NewCookieStore初始化一个store，同时传入一个secret key用来对session进行认证。
在Handler中，调用store.Get()获取一个已经存在的session或（如果不存在）创建一个新的。
设置sesssion.Values中的值，session.Values是map[interface{}]interface{}类型。
调用session.Save()将session保存到响应中

作者：kingeasternsun
链接：https://www.jianshu.com/p/91d457a85b36
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。`)
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}
