package middleware

import (
	"net/http"

	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/gjwt"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/util"
)

func AuthJwt() Adapter {
	//s1 := logg.Start()
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if gjwt.CheckJwt(w, r) {
				h.ServeHTTP(w, r)
			} else {
				msgjson := `{"status":"error","message":"Erro em seu token!"}`
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(msgjson))
				//io.WriteString(w, msgjson)
				//logg.Show(r.URL.Path, strings.ToUpper(r.Method), "error", s1)
			}
		})
	}
}

func AuthJwtNot(endpoints []string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if RegexUrlQuery(endpoints, r) {
				h.ServeHTTP(w, r)
				return
			}
			if gjwt.CheckJwt(w, r) {
				h.ServeHTTP(w, r)
			} else {
				msgjson := `{"status":"error","message":"token não está mais válido!"}`
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(msgjson))
				return
			}
		})
	}
}

func RegexUrlQuery(endpoints []string, r *http.Request) bool {
	///////////////////////////////////
	// to create routes dynamically
	// with regex for querys use in endpoints
	endpoint, _ := util.EndpointRegex(r)
	if len(endpoints) > 0 {
		for _, e := range endpoints {
			if e == endpoint {
				return true
			}
		}
	}
	return false
}
