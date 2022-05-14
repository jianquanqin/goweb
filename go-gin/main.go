package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//请求和响应在同一个函数中
	http.HandleFunc("/hello", sayhello)

	//监听端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("serve failed", err)
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {

	//从本地文件中读取内容
	b, _ := ioutil.ReadFile("./hello.html")
	//把读到的内容写到文件
	_, _ = fmt.Fprintln(w, string(b))
}
