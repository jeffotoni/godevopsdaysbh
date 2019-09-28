package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

// write bufio to optimization
var writer *bufio.Writer

func Println(text string) {
	writer = bufio.NewWriter(os.Stdout)
	writer.WriteString("\n")
	writer.WriteString(text)
	writer.Flush()
}

func main() {

	mux := http.NewServeMux()

	// diretorio fisico
	fs := http.FileServer(http.Dir("./web"))

	// mostra no browser localhost:8080/static
	mux.Handle("/", http.StripPrefix("/", fs))

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
