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
	http.HandleFunc("/hello", f1)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

func f1(w http.ResponseWriter, r *http.Request) {

	//三步走策略

	//1.定义模版
	//f.tmpl

	//2.解析模版
	//注意：模版路径要和二进制文件在同一目录下才能使用相对路径

	//第一种解析方式

	//新建一个模版对象f，同时调用上面的方法即进行解析模版
	template.New("f").ParseFiles("./f.tmpl")

	//第二种解析方式

	//t, err := template.ParseFiles("./hello.tmpl")
	//
	//if err != nil {
	//	fmt.Printf("Parse template failed, err:%v", err)
	//}

	//3，渲染模版

	////把解析的内容写给响应
	//name := "test"
	//err = t.Execute(w, name)
	//user1 := User{Name: "shiyivei", Gender: "male", Age: 18}
	////t.Execute(w, user1)
	//
	//if err != nil {
	//	fmt.Printf("Render template failed, err:%v", err)
	//}
}
