package main

import (
	"encoding/json"
	"net/http"

	"github.com/yuanyu90221/go-start/todo_server/controller"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
	// 		data := views.Response{
	// 			Code: http.StatusOK,
	// 			Body: "pong",
	// 		}
	// 		json.NewEncoder(w).Encode(data)
	// 	}
	// })
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
