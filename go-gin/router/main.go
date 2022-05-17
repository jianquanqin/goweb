package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//一个请求的路径和一个请求的方法对应一个handler
	//常见的有四个请求方法
	//获取信息
	//r.GET("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "GET",
	//	})
	//})
	//
	////创建信息，如注册成为会员
	//r.POST("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "POST",
	//	})
	//})
	//
	////删除数据
	//r.DELETE("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "DELETE",
	//	})
	//})
	//
	////修改部分数据
	//r.PUT("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "PUT",
	//	})
	//})

	//一种方式汇总所有请求方式
	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case "POST":
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		case "PUT":
			c.JSON(http.StatusOK, gin.H{"method": "PUT"})
		case "DELETE":
			c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
		}
	})
	//无路由的处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "baidu.com"})
	})

	//有公共前缀的路径
	////视频的首页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	////视频的详情页1
	//r.GET("/video/xxx", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/xxx"})
	//})
	////视频的详情页2
	//r.GET("/video/yyy", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/yyy"})
	//})

	//提取公共前缀
	videoGroup := r.Group("/video")

	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})

		})
		videoGroup.GET("/xxx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xxx"})

		})
		videoGroup.GET("/yyy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/yyy"})

		})
	}

	//商城的首页和详情页
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})
	err := r.Run(":8001")
	if err != nil {
		fmt.Printf("run server failed,err:%v", err)
	}
}
