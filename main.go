package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("layout.html", "home.html")
		if err != nil {
			log.Panic("解析失败", err)
		}
		t.ExecuteTemplate(w, "layout", "Hello World")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("layout.html", "about.html")
		if err != nil {
			log.Panic("解析失败", err)
		}
		t.ExecuteTemplate(w, "layout", "")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("layout.html")
		if err != nil {
			log.Panic("解析失败", err)
		}
		err = t.ExecuteTemplate(w, "layout", "")
		log.Println(err)
	})

	http.ListenAndServe("localhost:8080", nil)
}
