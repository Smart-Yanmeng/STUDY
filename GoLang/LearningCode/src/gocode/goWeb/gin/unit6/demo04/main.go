package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogFormatterParams(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"[YORK] %s %s| %d |%s %s| %s |%s %s\n",
		params.TimeStamp.Format("2006-04-02 15:04:05"),
		params.StatusCodeColor(), params.StatusCode, params.ResetColor(),
		params.MethodColor(), params.Method, params.ResetColor(),
		params.Path,
	)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// router := gin.Default()
	router := gin.New()

	// router.Use(gin.LoggerWithFormatter(LogFormatterParams))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: LogFormatterParams}))

	router.GET("/index", func(c *gin.Context) {

	})

	router.Run(":80")
}
