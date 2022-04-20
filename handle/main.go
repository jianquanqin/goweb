package main

import "net/http"

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func main() {

	mh := myHandler{}

	//第一种方式
	//http.ListenAndServe("localhost:8080", nil)
	//第二种方式：如下等同于上面的过程，但是分开写稍微更加灵活一点

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
}
