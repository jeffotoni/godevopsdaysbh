package main

import (
	"net/http"
)

func main() {
	println("Server Run... 8181")
	http.ListenAndServe(":8181", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"msg":"success"}`))
			return
		}))
}
