package controller

import (
	"encoding/json"
	"net/http"

	// "github.com/yuanyu90221/go-start/todo_server/views"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	return mux
}