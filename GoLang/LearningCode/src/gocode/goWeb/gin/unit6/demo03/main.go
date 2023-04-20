package main

import "github.com/gin-gonic/gin"

func main() {
	// 环境切换，控制台不输出
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {

	})

	router.Run(":80")
}
