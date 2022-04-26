package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("layout.html", "home.html")
		t.ExecuteTemplate(w, "layout.html", "Hello World")
	})
	http.ListenAndServe("localhost:8080", nil)
}
