// Go Api server
// @jeffotoni

package middleware

import (
	"net/http"
	"strings"
)

// CustomHeaders adds cache and basic HTTP headers to response
func CustomHeaders() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Server", "ApiProduto")
			w.Header().Add("Content-Type", "application/json")
			//w.Header().Add(runtime.HeaderContentType, "application/json")

			if strings.Contains(r.URL.RawQuery, "seed") {
				//deterministic calls can be cached
				w.Header().Add("Expires", "36000")
			} else {
				w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
				w.Header().Add("Pragma", "no-cache")
				w.Header().Add("Expires", "0")
			}

			h.ServeHTTP(w, r)
		})
	}
}
