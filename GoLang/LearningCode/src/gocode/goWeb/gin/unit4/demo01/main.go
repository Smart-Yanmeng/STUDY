package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()

	// 上传单个文件
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		// file.Size() 方法返回的结果的单位是字节
		fmt.Println(float64(file.Size)/1024, "KB")

		c.SaveUploadedFile(file, "./goWeb/gin/unit4/files/"+file.Filename)

		c.JSON(200, gin.H{"msg": "上传成功"})
	})

	// 上传单个文件
	router.POST("/read", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// 读取文件中的数据，返回文件对象
		readerFile, _ := file.Open()

		//data, _ := io.ReadAll(readerFile)
		//fmt.Println(string(data))

		writerFile, _ := os.Create("./goWeb/gin/unit4/files/" + file.Filename)
		defer writerFile.Close()
		n, _ := io.Copy(writerFile, readerFile)

		fmt.Println(n)

		c.JSON(200, gin.H{"msg": "读取成功"})
	})

	// 上传多个文件
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()

		files := form.File["upload[]"]

		for _, file := range files {
			c.SaveUploadedFile(file, "./goWeb/gin/unit4/files/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传 %d 个文件", len(files))})
	})

	router.Run(":80")
}
