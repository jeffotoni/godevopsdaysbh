// Go Api server
// @jeffotoni

package main

import (
	"fmt"
	"github.com/jeffotoni/godevopsdasybh/app.product/controller/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/login", handler.Login)
	http.Handle("/ping", http.HandlerFunc(handler.Ping))

	fmt.Println("Run Server: 8080")
	http.ListenAndServe(":8080", nil)
}
