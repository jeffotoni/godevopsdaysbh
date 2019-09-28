// Go Api server
// @jeffotoni

package middleware

import (
    "log"
    "net/http"
    "os"
    "time"
)

func Logger(name string) Adapter {

    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            h.ServeHTTP(w, r)
            //nao mostrar em ambiente de producao
            if os.Getenv("ENV_AMBI") != "PROD" {
                log.Printf(
                    "\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s\033[0m",
                    r.Method,
                    r.RequestURI,
                    name,
                    time.Since(start),
                )
            }
        })
    }
}
