package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name       string `json:"name"`        // 用户名
	Age        int    `json:"age"`         // 年龄
	Password   string `json:"password"`    // 密码
	RePassword string `json:"re_password"` // 确认密码
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var user UserInfo

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})

			return
		}
		c.JSON(200, gin.H{"data": user})
	})

	router.Run(":80")
}
