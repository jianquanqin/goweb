package main

import "net/http"

func main() {

	//1.注册一个函数，让其可以对web请求进行响应
	//处理web请求，我们一般用http包的下列函数，第一个参数相当于路由地址，"/"是根地址，表示响应所有的请求
	//第二个参数是个函数，一个是接口：用来写响应，一个是结构体的指针：包括传入请求的所有信息
	//这个函数是个回调函数，当请求到达时，函数就会执行（函数本身是参数，执行时，函数参数开始在最外层的函数体中发生一系列的行为）
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//接口w有write方法，所以可以调用
		w.Write([]byte("hello world"))
	})
	//2.设置web服务器，启动这个server
	//监听访问localhost:8080的请求，第二个参数默认调用Handler接口里面的方法（可以理解为一个路由器）
	//意思就是收到请求后，将所访问的路径分别匹配
	http.ListenAndServe("localhost:8080", nil) //DefaultServeMux
}
