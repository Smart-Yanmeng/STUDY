package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf(
			"[ York ] %s %s %s %d \n",
			httpMethod,
			absolutePath,
			handlerName,
			nuHandlers)
	}

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {})

	for _, info := range router.Routes() {
		fmt.Println(info.Path, info.Method, info.Handler)
	}

	router.Run(":80")
}
