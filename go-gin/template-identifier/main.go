package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {

	//定义模版
	//自定义函数
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{"safe": func(str string) template.HTML {
		return template.HTML(str)
	},
	//解析模版
	//并且使用链式结构解析
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//渲染模版
	//输入一段非法内容
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://liwenzhou.com'>liwenzhou的博客</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Printf("render template failed,err:%v\n", err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版：index.tmpl
	//解析模版
	//使用Delims()重新定义模版引擎标识符
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed,err:%v\n", err)
		return
	}
	//渲染模版
	msg := "小王子"
	err = t.Execute(w, msg)

	if err != nil {
		fmt.Printf("Render templates failed,err:%v\n", err)
		return
	}
}
