package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {

		//c.JSON(http.StatusOK, "ok")
		//第一种方式
		name := c.Query("query")
		age := c.Query("age")
		//第二种方式
		//name := c.DefaultQuery("query", "somebody")//如果制定了key就返回key对应的指定内容，否则返回默认值
		//第三种方式
		//name, ok := c.GetQuery("query") //和上述一致的意思
		//if !ok {
		//	name = "somebody"
		//}
		//返回,可以查询多个值
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":8080")
}
