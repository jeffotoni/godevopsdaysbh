package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func Ping(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method: ", r.Method)

	if strings.ToUpper(r.Method) == "POST" {

		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST...\n"))
}
