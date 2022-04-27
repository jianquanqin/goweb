package controller

import (
	"html/template"
	"log"
	"net/http"
)

func registerHomeRouts() {
	http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}

	t, err := template.ParseFiles("layout.html", "home.html")
	if err != nil {
		log.Panic("解析失败", err)
	}
	t.ExecuteTemplate(w, "layout", "Hello World")
}
