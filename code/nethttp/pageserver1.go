// Go Api server
// @jeffotoni

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Login struct {
	Title      string
	MsgError   string
	IfLabelone string
	Labelone   string
}

func HandlerLoginHtml(w http.ResponseWriter, r *http.Request) {
	// template lendo HTML

	tmpl := template.Must(template.ParseFiles("web/login.html"))

	login := Login{
		MsgError:   "",
		IfLabelone: "",
		Title:      "Workshop2.0",
		Labelone:   "logar!",
	}

	tmpl.Execute(w, login)
}

func HandlerAdminHtml(w http.ResponseWriter, r *http.Request) {

}

func HandlerAuth(w http.ResponseWriter, r *http.Request) {
	json := `{"status":"ok", "message":"tudo ocorreu bem na Auth..."}`
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func main() {

	mux := http.NewServeMux()

	// retorna HTML
	mux.HandleFunc("/login", HandlerLoginHtml)

	// auth
	mux.HandleFunc("/v1/api/auth", HandlerAuth)

	// fisico
	fs := http.FileServer(http.Dir("./web"))

	// vitual
	mux.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Server Run: 8085")
	http.ListenAndServe(":8085", mux)
}
