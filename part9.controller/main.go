package main

import (
	"github.com/solenovex/web/part9.controller/controller"
	"net/http"
)

func main() {

	controller.RegisterRouts()
	http.ListenAndServe("localhost:8080", nil)
}
