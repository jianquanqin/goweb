package main

import (
	"encoding/json"
	"github.com/solenovex/web/part11.midddleWare/middleware"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}
		time.Sleep(4 * time.Second)

		enc := json.NewEncoder(w)
		enc.Encode(c)
	})
	http.ListenAndServe("localhost:8080", &middleware.TimeoutMiddleware{Next: new(middleware.AuthMiddleware)})
}
