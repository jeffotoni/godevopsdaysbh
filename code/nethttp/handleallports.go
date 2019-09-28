// Go Api server
// @jeffotoni

package main

import (
	"fmt"
	//"log"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Curso golang porta 8082... \n")
}

func main() {

	fmt.Println("Server Run:8082")
	go func() {
		http.ListenAndServe(":8082", http.HandlerFunc(Index))
	}()

	go func() {
		http.ListenAndServe(":8083", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Curso golang 8083... \n")
		}))
	}()

	http.ListenAndServe(":8084", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Curso golang 8084... \n")
	}))
}
