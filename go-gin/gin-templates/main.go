package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	//创建一个gin模版引擎路由
	r := gin.Default()
	//如果有静态文件的话需要先加载
	r.Static("/xxx", "./statics")
	//在gin框架中给模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML { //定义匿名函数
			return template.HTML(str)
		},
	})
	//解析模版
	//r.LoadHTMLFiles("./templates/index.tmpl", "templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")

	//创建并处理GET请求
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //模版渲染
			"title": "posts/index.tmpl",
		})
	})
	//创建并处理GET请求
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //模版渲染
			"title": "<a href='liwenzhou.com'>liwenzhou的博客</a>",
		})
	})

	//返回从网上下载的模版
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.Run(":8080") //启动服务
}
