package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//其实这里只是定义有几级别目录，可以通过路径传参
	r.GET("/user/:name/:age", func(c *gin.Context) {

		//获取路径参数
		name := c.Param("name")
		age := c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")

		c.JSON(http.StatusOK, gin.H{
			"year": year,
			"age":  month,
		})
	})
	r.Run(":8080")
}
