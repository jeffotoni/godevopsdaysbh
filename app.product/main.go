// Go Api server
// @jeffotoni

package main

import (
	cf "github.com/jeffotoni/godevopsdasybh/app.product/config"
	"github.com/jeffotoni/godevopsdasybh/app.product/controller/handler"
	mw "github.com/jeffotoni/godevopsdasybh/app.product/controller/middleware"
	"log"
	"net/http"
)

func main() {

	// simple
	// http.HandleFunc("/login", handler.Login)
	// http.Handle("/ping", http.HandlerFunc(handler.Ping))
	// fmt.Println("Run Server: 8080")
	// http.ListenAndServe(":8080", nil)
	////////////////////////////////////////////////

	// DefaultServeMux
	mux := http.NewServeMux()

	// POST handler /login
	// Authorization: Basic key:keypass
	handlerApiLogin := http.HandlerFunc(handler.Login)

	// generate login jwt
	// handler login
	mux.Handle("/login", mw.Use(handlerApiLogin,
		mw.CustomHeaders(),
		mw.MaxClients(cf.MaxClients),
		mw.GtokenJwt(),
		mw.Logger("auth"),
	))

	// POST handler /ping
	handlerApiPing := http.HandlerFunc(handler.Ping)

	mux.Handle(cf.Endpoint().Ping, mw.Use(handlerApiPing,
		mw.CustomHeaders(),
		mw.Gzip(),
		mw.MaxClients(cf.MaxClients),
		mw.Logger("ping"),
	))

	server := &http.Server{
		Addr:    cf.HOST_CONFIG,
		Handler: mux,
		//Handler: middpromet,
		//ReadTimeout:  time.Millisecond * 600,
		//WriteTimeout: time.Millisecond * 400,
		//ReadTimeout:  cf.ReadTimeout,
		//WriteTimeout: cf.WriteTimeout,
		// IdleTimeout:    1000 * time.Millisecond,
		//MaxHeaderBytes: cf.MaxHeaderBytes,
	}

	log.Println("Run Server: ", cf.PORT_SERVER)
	log.Fatal(server.ListenAndServe())
}
