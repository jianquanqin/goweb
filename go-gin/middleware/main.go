package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//r := gin.Default() //默认使用了Logger()和Recovery这两个中间件
	r := gin.New()
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件m1,m2

	//请求先走m1再走中间件
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	////路由组1
	//xx1Group := r.Group("/xx1", authMiddleware(true))
	//{
	//	xx1Group.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"msg": "xx1Group",
	//		})
	//	})
	//}
	//
	////路由组2
	//xx2Group := r.Group("/xx2")
	//xx2Group.Use(authMiddleware(true))
	//{
	//	xx2Group.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"msg": "xx2Group",
	//		})
	//	})
	//}

	r.Run(":8082")
}

//把HandlerFunc提取出来

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}

//定义一个中间件

func m1(c *gin.Context) {
	//fmt.Println("m1")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "m1",
	//})
	fmt.Println("m1 in...")
	start := time.Now()
	c.Next() //调用后续处理的函数
	cost := time.Since(start)
	fmt.Println(cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")

	//在上下文中设置值，跨中间件
	c.Set("name", "shiyivei")
	//c.Next()  //调用后续处理的函数
	//c.Abort() //阻止调用后续处理的函数
	fmt.Println("m2 out...")
}

//一个常见的中间件模版
//func authMiddleware(x *gin.Context)  {
//是否登录判断
//if是登录用户
//c.next()
//else
//c.Abort()
//}

//但是更常写成闭包的形式
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其他的准备工作
	return func(c *gin.Context) {
		if doCheck {
			//是否登录判断
			//if是登录用户
			c.Next()
			//else
			//c.Abort()
		} else {
			c.Abort()
		}
	}
}
