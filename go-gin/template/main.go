package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	//2.解析模版
	//注意：模版路径要和二进制文件在同一目录下才能使用相对路径
	t, err := template.ParseFiles("./hello.tmpl")

	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
	}
	//3，渲染模版

	////把解析的内容写给响应
	//name := "test"
	//err = t.Execute(w, name)
	user1 := User{Name: "shiyivei", Gender: "male", Age: 18}
	t.Execute(w, user1)

	if err != nil {
		fmt.Printf("Render template failed, err:%v", err)
	}
}
