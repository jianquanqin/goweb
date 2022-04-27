package controller

import (
	"html/template"
	"log"
	"net/http"
)

func registerAboutRouts() {
	http.HandleFunc("/about", handleAbout)
}

func handleAbout(w http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("layout.html", "about.html")
	if err != nil {
		log.Panic("解析失败", err)
	}
	t.ExecuteTemplate(w, "layout", "")
}
