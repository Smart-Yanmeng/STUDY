package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func _json(c *gin.Context) {
	/* json 响应结构体 */
	type User struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		// 阻止对 Password 的 json 转换
		Password string `json:"-"`
	}
	// user := User{"York", 18, "123456"}
	// c.JSON(200, user)

	/* json 响应 map */
	// userMap := map[string]string{
	// 	 "user_name": "York",
	// 	 "age":       "18",
	// }
	// c.JSON(200, userMap)

	/* 直接响应 json */
	// c.JSON(200, gin.H{
	//	 "user_name": "York",
	//	 "age":       18,
	// })
}

func _xml(c *gin.Context) {
	c.XML(200, gin.H{
		"user_name": "York",
		"age":       18,
	})
}

func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{
		"user_name": "York",
		"age":       18,
	})
}

func _html(c *gin.Context) {
	// gin.H 加载从后端传过去的数据
	c.HTML(200, "index.html", gin.H{
		"username": "York",
	})
}

func _redirect(c *gin.Context) {
	// 访问网址
	c.Redirect(302, "https://www.baidu.com")
	// 访问路由
	// c.Redirect(302, "/html")
}

func main() {
	router := gin.Default()

	// 加载模板目录下所有的文件
	router.LoadHTMLGlob("goWeb/gin/unit2/demo02/templates/*")
	// 在 golang 中没有相对文件路径，只有相对项目路径
	router.StaticFile("/girl/girl.png", "goWeb/gin/unit2/demo02/static/girl.png")
	// 配置单个文件
	// 网页请求这个静态目录的前缀，第二个参数是一个目录（前缀不能重复）
	router.StaticFS("/static", http.Dir("goWeb/gin/unit2/demo02/static/static"))

	// 响应字符串
	router.GET("/string", _string)

	// 响应 JSON
	router.GET("/json", _json)

	// 响应 XML
	router.GET("/xml", _xml)

	// 响应 YAML
	router.GET("/yaml", _yaml)

	// 响应 HTML
	router.GET("/html", _html)

	// 重定向
	router.GET("/redirect", _redirect)

	router.Run(":80")
}
