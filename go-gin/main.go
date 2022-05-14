package main

import "github.com/gin-gonic/gin"

//使用Gin框架

func main() {

	r := gin.Default() //返回默认的路由引擎

	//访问路径以及返回
	r.GET("/hello", sayhello)

	//启动
	//封装了http.ListenAndServe
	r.Run(":8080")
}

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Golang",
	})
}

//使用Go语言内置的代码创建web服务

//func main() {
//
//	//请求和响应在同一个函数中
//	http.HandleFunc("/hello", sayhello)
//
//	//监听端口
//	err := http.ListenAndServe(":8080", nil)
//
//	if err != nil {
//		fmt.Println("serve failed", err)
//	}
//}
//
//func sayhello(w http.ResponseWriter, r *http.Request) {
//
//	//从本地文件中读取内容
//	b, _ := ioutil.ReadFile("./hello.html")
//	//把读到的内容写到文件
//	_, _ = fmt.Fprintln(w, string(b))
//}
