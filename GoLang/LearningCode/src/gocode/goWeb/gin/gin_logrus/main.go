package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gocode/goWeb/gin/gin_logrus/log"
	"gocode/goWeb/gin/gin_logrus/middleware"
)

func main() {
	log.InitFile("goWeb/gin/gin_logrus/logs", "server")

	router := gin.New()
	router.Use(middleware.LogMiddleware())

	router.GET("/", func(c *gin.Context) {
		logrus.Info("INFO")

		c.JSON(200, gin.H{"msg": "Hello"})
	})

	router.Run(":80")
}
