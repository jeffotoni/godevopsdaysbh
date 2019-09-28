// Go Api server
// @jeffotoni

package middleware

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

//write in gzip and Header() from http
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Gzip functionality if the clients accepts it
func Gzip() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			algSupp := r.Header.Get("Accept-Encoding")
			supportGzip := strings.Contains(algSupp, "gzip")

			if supportGzip {
				w.Header().Set("Content-Encoding", "gzip")
				gz := gzip.NewWriter(w)

				defer func() {
					err := gz.Close()
					if err != nil {
						log.Printf("error closing gzip: %+v\n", err)
					}
				}()

				gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
				h.ServeHTTP(gzr, r)
				return
			}

			//fallback
			h.ServeHTTP(w, r)
		})
	}
}
