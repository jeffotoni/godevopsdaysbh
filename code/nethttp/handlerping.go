// Go Api server
// @jeffotoni

package main

import (
	"log"
	"net/http"
)

func main() {

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nWorkhop2.0 Golang HandleFunc!"))
	}

	// handleFunc
	http.HandleFunc("/v1/api/ping", pingHandler)
	http.HandleFunc("/v1/api/ping2", pingHandler)
	http.HandleFunc("/v1/api/ping3", pingHandler)

	// show run server
	log.Printf("\nServer run 8085\n")

	// Listen
	log.Fatal(http.ListenAndServe(":8085", nil))
}
