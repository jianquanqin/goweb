package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		////获取的赋值给变量
		//username := c.Query("username")
		//password := c.Query("password")
		//
		////把变量存起来
		//user := UserInfo{username: username, password: password}
		//fmt.Printf("%#v\n", user)

		var user UserInfo
		err := c.ShouldBind(&user) //使用指针才能改变原来的值
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			//返回一个状态
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})

	r.POST("/form", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user) //使用指针才能改变原来的值
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			//返回一个状态
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user) //使用指针才能改变原来的值
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			//返回一个状态
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":8080")
}
