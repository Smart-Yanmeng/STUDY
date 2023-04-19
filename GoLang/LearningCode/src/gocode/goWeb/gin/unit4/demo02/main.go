package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {
		// 表示是文件流，唤起浏览器下载
		c.Header("Content-Type", "application/octet-stream")

		// 用来指定下载下来的文件名
		c.Header("Content-Disposition", "attachment; filename="+"test01.png")

		// 表示传输过程中的编码形式，乱码的问题可能就是因为他
		c.Header("Content-Transfer-Encoding", "binary")

		c.File("./goWeb/gin/unit4/files/test01.png")
	})

	router.Run(":80")
}
