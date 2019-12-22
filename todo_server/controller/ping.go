package controller

import (
	"encoding/json"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(data)
		}
	}
}