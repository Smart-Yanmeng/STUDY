package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()

	// 请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		// 首字母大小写不区分，单词与单词之间用 "-" 连接
		// 用于获取一个请求头
		fmt.Println("-----------------------")

		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.GetHeader("user-Agent"))
		fmt.Println(c.GetHeader("user-agent"))

		// 用于获取一个请求头 map 集合
		// Header 是一个普通的 map[string][]string
		// 如果是使用 Get 方法或者是 GetHeader，那么可以不用区分大小写，并且返回第一个 value
		// 如果是用 map 的取值方式，请注意大小写问题
		fmt.Println("-----------------------")

		fmt.Println(c.Request.Header)
		fmt.Println(c.Request.Header.Get("User-Agent"))
		fmt.Println(c.Request.Header["User-Agent"])

		fmt.Println("-----------------------")

		// 自定义的请求头，用 Get 方法也是不用区分大小写的
		fmt.Println(c.Request.Header.Get("token"))

		c.JSON(200, gin.H{"msg": "success"})
	})

	// 爬虫和用户区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		// 用正则去匹配
		// 字符串的包含匹配
		if strings.Contains(userAgent, "python") {
			// 爬虫来了
			c.JSON(0, gin.H{"data": "这是响应给爬虫的数据"})

			return
		}
		c.JSON(0, gin.H{"data": "这是响应给用户的数据"})
	})

	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "hdakjdjkadhkajdhwahd")
		c.Header("Content-Type", "application/text")

		c.JSON(200, gin.H{"data": "看看响应头"})
	})

	router.Run(":80")
}
