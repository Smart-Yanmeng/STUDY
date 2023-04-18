package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	// required - 不能为空，且不能没有这个字段
	// min/max/len - 针对字符串的长度，最小/最大/等于字符数
	// contains - 必须包含固定字符
	// startswith/endswith - 以固定开始/结尾
	Name string `json:"name" binding:"required,min=4,max=8,len=6,contains=o,startswith=Y,endswith=6"` // 用户名
	// gte/lte - 针对数字大小，大于等于/小于等于数字
	Age      int    `json:"age" binding:"gte=18,lte=30"` // 年龄
	Password string `json:"password"`                    // 密码
	// eqfield - 针对同级字段，需要两个字段的值相同
	RePassword string `json:"re_password" binding:"eqfield=Password"` // 确认密码
	// oneof - 字段必须为其中之一
	Sex string `json:"sex" binding:"oneof=man woman"` // 性别
	// dive - 针对数组中每一个元素
	LikeList []string `json:"like_list" binding:"required,dive,startswith=like"` // 爱好
	// ip - 针对 IP 地址，必须为一个有效的 IP 地址
	IP string `json:"ip" binding:"ip"` // IP 地址
	// ip - 针对 Url 地址，必须为一个有效的 Url 地址
	Url string `json:"url" binding:"url"` // Url 地址
	// ip - 针对 Uri 地址，必须为一个有效的 Uri 地址
	Uri string `json:"uri" binding:"uri"` // Uri 地址 - 是 Url 地址的子集
	// date - 针对日期，必须为一个正确的日期格式
	Date string `json:"date" binding:"datetime=2006-01-02 15:04:05"`
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
