package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 拿到一个*gin.Engine
	r := gin.Default()

	// 获取 GET 请求
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	// 开启服务
	r.Run()
}
