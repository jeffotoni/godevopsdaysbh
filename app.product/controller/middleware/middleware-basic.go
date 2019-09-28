// Go Api server
// @jeffotoni

package middleware

import (
    "net/http"
)

func AutBasic() Adapter {
    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            var ok bool = true
            var jsonerr string = ""
            if ok {
                h.ServeHTTP(w, r)
            } else {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusUnauthorized)
                w.Write([]byte(jsonerr))
            }
        })
    }
}
