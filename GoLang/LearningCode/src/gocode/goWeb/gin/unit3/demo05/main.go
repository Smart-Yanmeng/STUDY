package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	// required - 不能为空，且不能没有这个字段
	// min/max/len - 针对字符串的长度，最小/最大/等于字符数
	Name string `json:"name" binding:"required,min=4,max=8,len=6"` // 用户名
	// gte/lte - 针对数字大小，大于等于/小于等于数字
	Age      int    `json:"age" binding:"gte=18,lte=30"` // 年龄
	Password string `json:"password"`                    // 密码
	// eqfield - 针对同级字段，需要两个字段的值相同
	RePassword string `json:"re_password" binding:"eqfield=Password"` // 确认密码
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
