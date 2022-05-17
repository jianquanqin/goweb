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
	http.HandleFunc("/templates", Template)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

func Template(w http.ResponseWriter, r *http.Request) {

	//定义
	//解析
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed, err:%v", err)
		return
	}
	//渲染
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render templates failed, err:%v", err)
		return
	}
}

func f1(w http.ResponseWriter, r *http.Request) {

	//当创建了一个自定义的函数时需要告诉（向模版引擎注册）模版引擎这个事件

	//1.定义模版 创建f.tmpl文件

	//2.解析模版
	//注意：模版路径要和二进制文件在同一目录下才能使用相对路径

	//第一种解析方式

	//新建一个模版对象f，同时调用上面的方法即进行解析模版
	//这也就是常用的链式操作
	//注意：创建的文件名称要和打开的文件名称一样
	t := template.New("f.tmpl")

	//定义模版函数
	//注意；自定义函数要么返回一个值，要么返回两个值，当返回了两个值时，第二个值必须是error类型
	customizeFunc := func(arg string) (string, error) {
		return arg + "聪明又可爱", nil
	}
	//k是调用时所指的名称
	t.Funcs(template.FuncMap{"Praise": customizeFunc})

	t, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed, err:%v", err)
	}

	//第二种解析方式

	//t, err := templates.ParseFiles("./hello.tmpl")
	//
	//if err != nil {
	//	fmt.Printf("Parse templates failed, err:%v", err)
	//}

	//3.渲染模版

	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render templates failed, err:%v", err)
	}

}
