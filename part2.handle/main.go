package main

import "net/http"

//先自定义一个类型
type helloHandler struct{}

//为它实现ServeHTTP方法后就会将它变成一个handler
func (mh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

//再自定义一个类型
type aboutHandler struct{}

//同样的，实现ServeHTTP方法
func (mh *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {

	mh := helloHandler{} //初始化变量
	a := aboutHandler{}  //初始化变量

	//第一种方式
	//http.ListenAndServe("localhost:8080", nil)
	//第二种方式：如下等同于上面的过程，但是分开写稍微更加灵活一点

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	//1.使用Handle向DefaultServeMux注册，此时再过来符合条件的请求就会经由DefaultServeMux分发给该注册过的Handle
	//其实就是把我们自定义的handler加入到DefaultServeMux分配的备选handler集合之中
	http.Handle("/hello", &mh) //注册
	http.Handle("/about", &a)  //同上也是注册

	//2.使用HandleFunc向DefaultServeMux注册
	//在签名中定义func
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home!"))
	})
	//在外面定义，直接传函数名
	http.HandleFunc("/welcome", welcome)

	//3.也可以使用http.HandlerFunc()将函数转换为为一个handler，其实就是将函数名称变为ServeHTTP
	//type HandlerFunc func(ResponseWriter, *Request)
	//
	//// ServeHTTP calls f(w, r).
	//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	//	f(w, r)
	//}
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
}
