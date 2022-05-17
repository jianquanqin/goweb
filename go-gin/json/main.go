package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//使用gin框架的步骤
	//1.引入
	r := gin.Default()

	//2.写访问路径和handler
	//方法一,使用map
	r.GET("/json", func(c *gin.Context) {

		//定义一个map，填入数据
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}

		//or使用内置的map进行填写
		data := gin.H{"name": "小王子", "message": "hello world", "age": 18}
		c.JSON(http.StatusOK, data)
	})

	//方法二，使用struct
	//type msg struct {
	//	Name    string
	//	Message string
	//	Age     int
	//}

	//使用tag标签可以把能导出的字段设置为自定义的格式，如把Name 显示成 name
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	//2.写访问路径和handler
	r.GET("/another_json", func(c *gin.Context) {

		data := msg{
			Name:    "大王子",
			Message: "Hello golang",
			Age:     19,
		}
		c.JSON(http.StatusOK, data)
	})
	//3.启动,写上端口号
	r.Run(":8080")
	//输出结果： {"name":"大王子","message":"Hello golang","age":19}
}
