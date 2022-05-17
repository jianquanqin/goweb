package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		//跳转百度
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//内部重定向
	r.GET("/a", func(c *gin.Context) {
		//跳转到b对应的路由处理函数
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "bbb",
		})
	})

	r.Run(":8080")
}
