package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询参数
func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	// 返回一个键值对
	fmt.Println(c.GetQuery("user"))
	// 拿到多个相同的查询参数
	fmt.Println(c.QueryArray("user"))
	// 拿到多个相同的查询参数
	fmt.Println(c.QueryMap("user"))
	// 如果用户没传值，就使用默认值
	fmt.Println(c.DefaultQuery("addr", "浙江省"))
}

// 动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
}

// 表单参数
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	// 如果用户没传值，就使用默认值
	fmt.Println(c.DefaultPostForm("addr", "浙江省"))
	form, err := c.MultipartForm() // 接受所有的 form 参数，包括文件
	fmt.Println(form, err)
}

// 自封装函数 - 解析 json 到结构体上的函数
func bindJson(c *gin.Context, obj any) (err error) {
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

// 原始参数
func _raw(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	err := bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	//body, _ := c.GetRawData()
	//fmt.Println(string(body))
	//
	//contentType := c.GetHeader("Content-Type")
	//fmt.Println(contentType)
	//
	//switch contentType {
	//case "application/json":
	//	// json 解析到结构体中
	//	type User struct {
	//		Name string `json:"name"`
	//		Age  int    `json:"age"`
	//	}
	//
	//	var user User
	//	err := json.Unmarshal(body, &user)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//
	//		return
	//	}
	//	fmt.Println(user)
	//}
}

func main() {
	router := gin.Default()

	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)

	router.Run(":8080")
}
