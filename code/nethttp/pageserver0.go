package main

import "net/http"

func main() {

	// http://localhost:8085/login.html

	// diretorio fisico
	fs := http.FileServer(http.Dir("web/"))
	// mostra no browser localhost:8080/static
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8085", nil)
}
