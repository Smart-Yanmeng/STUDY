package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1 in")

	// c.JSON(200, gin.H{"msg": "m1的响应"})
	// 停止后续的响应
	// c.Abort()

	// 继续后续的响应
	c.Next()

	fmt.Println("m1 out")
}

func index(c *gin.Context) {
	fmt.Println("index in")

	c.JSON(200, gin.H{"msg": "success"})
	c.Abort()
	c.Next()

	fmt.Println("index out")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in")

	c.Next()

	fmt.Println("m2 out")
}

func main() {
	router := gin.Default()

	// 可以响应多个中间件
	router.GET("/", m1, index, m2)

	router.Run(":80")
}
