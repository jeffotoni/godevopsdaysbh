package main

import (
	"github.com/gorilla/mux"
	. "handler"
	"log"
	"net/http"
	"pkg/rpc"
)

func main() {

	//server RPC
	go rpcx.Run()

	// server HTTP
	r := mux.NewRouter()

	r.HandleFunc("/user", UserGetAll).Methods("GET")

	server :=
		&http.Server{
			Addr:    ":8080",
			Handler: r,
		}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error ao subir server: %s", err)
	}
}
