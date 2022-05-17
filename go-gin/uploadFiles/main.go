package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()

	//先写一个用来上传文件的html页面
	r.LoadHTMLFiles("./index.html")

	//使用Get请求访问页面并让用户上传文件
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//用户提交之后使用post的请求将文件处理
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//filePath := fmt.Sprintf("./%s",f.Filename)
			filePath := path.Join("./", f.Filename) //写到本文件夹下
			err = c.SaveUploadedFile(f, filePath)
			if err != nil {
				fmt.Printf("save uploadfile failed, err:%v", err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":8080")
}
