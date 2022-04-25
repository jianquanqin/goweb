package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http.HandleFunc("/write", writerExample)
	//http.HandleFunc("/writeheader", writeHeaderExample)
	//http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //设置content-type
	post := &Post{
		User:    "Sau sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post) //将struct转换为它的json编码，实际上是一个byte slice
	w.Write(json)                 //使用Write方法返回回去
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "https://www.google.com") //第一个参数是key，第二个参数是value
	w.WriteHeader(302)
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no such service, try next door")
}

func writerExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head>Go Web</head> //head里面只写了内容
<body><h1>Hello World</h1></body> 
</html>`
	w.Write([]byte(str))
}
