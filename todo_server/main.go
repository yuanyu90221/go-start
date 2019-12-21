package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// fmt.Pirnt("test"
		}
	})
	http.ListenAndServe(":3000", nil)
}
