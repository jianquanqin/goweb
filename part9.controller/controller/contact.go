package controller

import (
	"html/template"
	"log"
	"net/http"
)

func registerContactRouts() {
	http.HandleFunc("/contact", handleContact)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("layout.html")
	if err != nil {
		log.Panic("解析失败", err)
	}
	err = t.ExecuteTemplate(w, "layout", "")
	log.Println(err)
}
