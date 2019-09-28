//go get github.com/rakyll/statik
//go mod init serverworkshop

package main

import (
	"bufio"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	"os"
	_ "statik" // TODO: Replace with the absolute import path
	"strings"
)

var writer *bufio.Writer

func Println(text string) {
	writer = bufio.NewWriter(os.Stdout)
	writer.WriteString("\n")
	writer.WriteString(text)
	writer.Flush()
}

func main() {

	mux := http.NewServeMux()

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// diretorio fisico
	fs := http.FileServer(statikFS)

	// mostra no browser localhost:8080/static
	mux.Handle("/", http.StripPrefix("/", DisabledFs(fs)))

	Println("Run Server")
	http.ListenAndServe(":8080", mux)
}

func DisabledFs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
