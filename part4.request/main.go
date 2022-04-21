package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	//1.#fragment
	//http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, r.URL.Fragment)
	//})

	//2.Header
	// http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
	// 	//fmt.Fprintln(w, r.URL.Fragment)
	// 	fmt.Fprintln(w, r.Header)
	// 	fmt.Fprintln(w, r.Header["Accept-Encoding"])
	// 	fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))

	// })

	//3.Body
	//http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
	//	length := r.ContentLength //获取内容的长度
	//	body := make([]byte, length)
	//	r.Body.Read(body) //获取body，  并将其读到上述bytes中
	//
	//	fmt.Fprintln(w, string(body))
	//})

	//4.URL
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query() //返回的是一个map
		fmt.Println(query)

		id := query["id"] //返回的是id []string
		log.Println(id)

		name := query.Get("name") //返回的是[]string中的第一个元素
		log.Println(name)
	})

	server.ListenAndServe()
}
