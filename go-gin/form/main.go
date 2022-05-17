package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//第一种方式
		//获取form表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//第二种方式
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***") //key取不到说明没有此键

		//第三种方法
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":8080")
}
