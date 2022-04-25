package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	//http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
	//	//http.ServeFile(w, r, "form"+r.URL.Path) 因为在表单里面写了地址了，所以不用再写访问路径
	//	//r.ParseForm() //先解析表单
	//	//r.ParseMultipartForm(1024) //此处指字节长度
	//
	//	// fmt.Fprintln(w, r.PostForm["first_name"]) //map[first_name:[yivei] last_name:[shi]]
	//	// fmt.Fprintln(w, r.Form["first_name"])     // [yivei Nick]
	//	// fmt.Fprintln(w, r.PostForm["first_name"]) //[yivei]
	//	//fmt.Fprintln(w, r.MultipartForm) // &{map[first_name:[yivei] last_name:[shi]] map[]}
	//	//fmt.Fprintln(w, r.FormValue("first_name")) //不用解析，使用FormValue直接调用 //Nick
	//
	//})
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

//写一个handler
func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(1024) //解析表单
	//
	//fileHandler := r.MultipartForm.File["uploaded"][0] //访问File字段
	//file, err := fileHandler.Open()                    //获取文件

	file, _, err := r.FormFile("uploaded") //该方法能够直接获取文件而不用解析（默认是获取slice的第一个值）

	if err != nil {
		data, err := ioutil.ReadAll(file) //读取文件
		if err == nil {
			fmt.Fprintln(w, string(data)) //打印
		}
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
