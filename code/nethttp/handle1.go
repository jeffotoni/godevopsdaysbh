// Go Api server
// @jeffotoni

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var sb strings.Builder

func Send(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebendo Json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)

		jsonError := `{"status":"error", 
		"msg":"` + err.Error() + `"}`

		W.Write([]byte(jsonError))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST...\n"))
}

func Ping(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method: ", r.Method)

	if strings.ToUpper(r.Method) == "POST" {

		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST...\n"))
}

func main() {

	Ping2 := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		json := `{"status":"ok", "msg":"sucesso"}`
		json += "\n"
		w.Write([]byte(json))
	}

	http.HandleFunc("/api/v1/ping1", Ping)
	http.Handle("/api/v1/ping2", http.HandlerFunc(Ping2))

	http.HandleFunc("/api/v1/ping3", Ping)
	http.HandleFunc("/api/v1/ping4", Ping)

	fmt.Println("Run Server: 8085")
	http.ListenAndServe(":8085", nil)
}
