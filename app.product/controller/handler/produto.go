package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Produto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebendo Json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)

		jsonError := `{"status":"error", 
		"msg":"` + err.Error() + `"}`

		w.Write([]byte(jsonError))
		return
	}

	fmt.Println(string(b))

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST...\n"))
}
