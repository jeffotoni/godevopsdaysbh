// Go Api server
// @jeffotoni

package middleware

import (
    "net/http"
)

var (
    CorsAllow          = []string{"", ""}
    CorsAllowedMethods = []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"}
    CorsAllowedHeaders = []string{"*"}
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Cors() Adapter {
    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            enableCors(&w)
            setupResponse(&w, r)
            if (*r).Method == "OPTIONS" {
                println("estou here...")
                return
            }
            h.ServeHTTP(w, r)
        })
    }
}
