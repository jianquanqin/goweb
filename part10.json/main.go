package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodPost:
	// 		dec := json.NewDecoder(r.Body)
	// 		company := Company{}
	// 		err := dec.Decode(&company)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			return
	// 		}
	// 		enc := json.NewEncoder(w)
	// 		err = enc.Encode(company)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			return
	// 		}
	// 	default:
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 	}
	// })
	jsonStr := `
	{
		"id" :123,
		"name" : "Google",
		"country": "USA"
	}`
	c := Company{}
	_ = json.Unmarshal([]byte(jsonStr), &c)
	fmt.Println(c)

	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	bytes1, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(bytes1))

	http.ListenAndServe("localhost:8080", nil)
}
