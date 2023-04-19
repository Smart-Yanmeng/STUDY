package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Data any
	Msg  string
}

func UserListView(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var userList []UserInfo = []UserInfo{
		{"York", 18},
		{"Nan", 18},
		{"Jie", 18},
	}

	c.JSON(200, Response{0, userList, "success"})
}

func Middleware(msg string) gin.HandlerFunc {
	// 闭包
	fmt.Println("这里的代码立马执行")

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "123456" {
			c.Next()

			return
		}
		c.JSON(200, Response{1001, nil, msg})

		c.Abort()
	}
}

func UserRouterInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager").Use(Middleware("用户验证失败"))
	{
		userManager.GET("/users", UserListView)
	}
}

func main() {
	// Default() 中有 Logger 和 Recovery 的中间件
	// New() 中不包含任何的中间件
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("api")

	api.GET("/login", Middleware("登录验证失败"), func(c *gin.Context) {
		panic("手动报错")

		c.JSON(200, gin.H{"data": "123456"})
	})

	UserRouterInit(api)

	router.Run(":80")
}
