package main

import (
	"github.com/solenovex/web/part9.controller/controller"
	"net/http"
)

func main() {

	controller.RegisterRouts()
	http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil)
}
