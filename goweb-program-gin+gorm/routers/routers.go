package routers

import (
	"github.com/gin-gonic/gin"
	"gorm/program/controller"
)

func SetupRouter() *gin.Engine {
	//4.创建路由开始对所有请求进行处理
	r := gin.Default()

	//加载静态文件
	r.Static("/static", "static")
	//加载模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//根据用户请求对数据库进行操作
	v1Group := r.Group("v1")

	{
		//添加数据
		v1Group.POST("/todo", controller.CreateHandler)
		//查看所有待办事项
		v1Group.GET("/todo", controller.ViewHandler)
		//更新某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateHandler)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteHandler)
	}
	return r
}
