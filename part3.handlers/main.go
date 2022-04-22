package main

import "net/http"

func main() {

	//第一种方式
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "form"+r.URL.Path)
	//})
	//
	//http.ListenAndServe(":8080", nil)
	//第二种方式
	http.ListenAndServe(":8080", http.FileServer(http.Dir("index")))
}
