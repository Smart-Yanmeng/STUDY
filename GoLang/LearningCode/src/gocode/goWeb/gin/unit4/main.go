package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		
	})

	router.Run(":80")
}
