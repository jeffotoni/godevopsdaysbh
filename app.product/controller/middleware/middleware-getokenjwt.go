// Go Api server
// @jeffotoni
// 2019-01-04

package middleware

import (
	"net/http"

	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/gjwt"
)

func GtokenJwt() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ok, msgjson, _ := gjwt.CheckBasic(w, r)
			if ok {
				h.ServeHTTP(w, r)
			} else {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(msgjson))
			}
		})
	}
}
