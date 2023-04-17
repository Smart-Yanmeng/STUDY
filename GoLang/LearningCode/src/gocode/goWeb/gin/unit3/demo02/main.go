package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	fmt.Println(string(body))

	contentType := c.GetHeader("Content-Type")
	fmt.Println(contentType)

	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())

			return err
		}
	}
	return nil
}

// 文章列表页面
func _getList(c *gin.Context) {
	articleList := []ArticleModel{
		{"Go", "go语言入门"},
		{"Python", "python语言入门"},
		{"Java", "java语言入门"},
	}

	c.JSON(200, Response{0, articleList, "success"})
}

// 文章详情
func _getDetail(c *gin.Context) {
	// 获取 param 中的 id
	fmt.Println(c.Param("id"))
	article := ArticleModel{"Go", "go语言入门"}

	c.JSON(200, Response{0, article, "success"})
}

// 创建文章
func _create(c *gin.Context) {
	// 接受前端纯进来的 json 数据
	var article ArticleModel

	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)

		return
	}

	c.JSON(200, Response{0, article, "success"})
}

// 编辑文章
func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var article ArticleModel

	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)

		return
	}

	c.JSON(200, Response{0, article, "success"})
}

// 删除文章
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))

	c.JSON(200, Response{0, map[string]string{}, "success"})
}

func main() {
	router := gin.Default()

	router.GET("/articles", _getList)       // 文章列表
	router.GET("/articles/:id", _getDetail) // 文章详情
	router.POST("/articles", _create)       // 添加文章
	router.PUT("/articles/:id", _update)    // 编辑文章
	router.DELETE("/articles/:id", _delete) // 删除文章

	router.Run(":80")
}
