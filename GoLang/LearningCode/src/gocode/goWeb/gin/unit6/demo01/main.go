package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func indexName(c *gin.Context) {

}

func main() {
	// 改变日志为文件输出
	file, _ := os.Create("gin.log")

	// os.Stdout 表示终端输出日志
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()

	// 匿名函数的日志输出
	// handlers 表示中间件的个数，默认带两个中间件
	// [GIN-debug] GET    /index                    --> main.main.func1 (3 handlers)
	router.GET("/index", func(c *gin.Context) {})
	router.POST("/users", func(c *gin.Context) {})
	router.POST("/articles", func(c *gin.Context) {})

	// 具名函数的日志输出
	// [GIN-debug] GET    /indexName                --> main.indexName (3 handlers)
	router.GET("/indexName", indexName)

	api := router.Group("api")

	api.DELETE("/articles/:id", func(c *gin.Context) {})

	router.Run()
}
