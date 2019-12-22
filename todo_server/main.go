package main

import (
	// "encoding/json"
	"net/http"
	"fmt"
	"github.com/yuanyu90221/go-start/todo_server/controller"
)

func main() {
	mux := controller.Register()
	fmt.Println("test");
	http.ListenAndServe(":3000", mux)
}
