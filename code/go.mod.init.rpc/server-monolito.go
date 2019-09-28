package main

import (
	"github.com/gorilla/mux"
	. "handler"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	//r.HandleFunc("/auht", Auth).Methods("POST")
	r.HandleFunc("/user", UserGetAll).Methods("GET")
	// r.HandleFunc("/user/{uuid}", UserGet).Methods("GET")
	// r.HandleFunc("/user/{uuid}", UserDel).Methods("DELETE")
	// r.HandleFunc("/user/{uuid}", UserPut).Methods("PUT")

	server :=
		&http.Server{
			Addr:    ":8080",
			Handler: r,
		}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error ao subir server: %s", err)
	}
}
