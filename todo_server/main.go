package main

import (
	// "encoding/json"
	"net/http"

	"github.com/yuanyu90221/go-start/todo_server/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
