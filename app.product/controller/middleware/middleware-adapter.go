// Go Api server
// @jeffotoni

package middleware

import (
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func Use(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
