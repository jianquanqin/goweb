package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	http.HandleFunc("/index1", index1)
	http.HandleFunc("/home1", home1)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版：index1.tmpl
	//解析模版
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed,err:%v\n", err)
		return
	}
	//渲染模版
	msg := "小王子"
	t.Execute(w, msg)

}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模版：home1.tmpl
	//解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed,err:%v\n", err)
		return
	}
	//渲染模版
	msg := "小王子"
	t.Execute(w, msg)
}

func index1(w http.ResponseWriter, r *http.Request) {
	//定义模版：index1.tmpl
	//解析模版,注意顺序
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index1.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed,err:%v\n", err)
		return
	}
	//渲染模版
	msg := "小王子"
	t.ExecuteTemplate(w, "index1.tmpl", msg)

}

func home1(w http.ResponseWriter, r *http.Request) {
	//定义模版：home1.tmpl
	//解析模版,注意顺序
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home1.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed,err:%v\n", err)
		return
	}
	//渲染模版
	msg := "小王子"
	t.ExecuteTemplate(w, "home1.tmpl", msg)

}
