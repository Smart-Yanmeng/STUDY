package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func m1(c *gin.Context) {
	fmt.Println("m1 in")
	//c.JSON(200, gin.H{"msg": "哦豁，你的响应被我吃掉了"})
	//c.Abort()
	//c.Next()
	//fmt.Println("m1 out")

	c.Set("name", "York")
	c.Set("user", User{
		Name: "York",
		Age:  18,
	})
}

//func m2(c *gin.Context) {
//	fmt.Println("m2 in")
//c.Next()
//fmt.Println("m2 out")
//}

func main() {
	router := gin.Default()

	// 全局注册中间件
	// router.Use(m1, m2)
	router.Use(m1)

	router.GET("/m1", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println(name)

		// 断言
		_user, _ := c.Get("user")
		user := _user.(User)
		fmt.Println("name =", user.Name)
		fmt.Println("age =", user.Age)

		//fmt.Println("index in")
		c.JSON(200, gin.H{"msg": "m1"})
		//c.Next()
		//fmt.Println("index out")
	})
	//router.GET("/m2", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"msg": "m2"})
	//})

	router.Run(":80")
}
