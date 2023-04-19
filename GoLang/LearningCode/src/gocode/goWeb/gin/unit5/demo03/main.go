package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int
	Data any
	Msg  string
}

func UserListView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"York", 18},
		{"Nan", 18},
		{"Jie", 18},
	}

	c.JSON(200, Response{0, userList, "success"})
}

func ArticleListView(c *gin.Context) {
	var articleList []ArticleInfo = []ArticleInfo{
		{"Python", "Python"},
		{"GoLang", "GoLang"},
		{"Java", "Java"},
	}

	c.JSON(200, Response{0, articleList, "success"})
}

func UserRouterInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager")
	{
		userManager.GET("/users", UserListView)
		userManager.POST("/users", UserListView)
	}
}

func ArticleRouterInit(router *gin.RouterGroup) {
	articleManager := router.Group("article_manager")
	{
		articleManager.GET("/article", ArticleListView)
	}
}

func main() {
	router := gin.Default()

	// 一个路由组，请求的时候需要加上 "/api"
	api := router.Group("api")

	// 路由嵌套
	//userManager := api.Group("user_manager")
	//{
	//	userManager.GET("/users", UserListView)
	//	userManager.POST("/users", UserListView)
	//}
	//articleManager := api.Group("article_manager")
	//{
	//	articleManager.GET("/article", ArticleListView)
	//}

	UserRouterInit(api)
	ArticleRouterInit(api)

	router.Run(":80")
}
