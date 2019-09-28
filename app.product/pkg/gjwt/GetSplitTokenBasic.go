package gjwt

import (
	"net/http"
	"strings"
)

//takes the basic token and returns to work
func GetSplitTokenBasic(w http.ResponseWriter, r *http.Request) string {
	var Authorization string
	Authorization = r.Header.Get("Authorization")
	if Authorization == "" {
		Authorization = r.Header.Get("authorization")
	}
	// browsers
	if Authorization == "" {
		Authorization = r.Header.Get("Access-Control-Allow-Origin")
	}

	if Authorization != "" {
		auth := strings.SplitN(Authorization, " ", 2)
		if len(auth) != 2 || strings.TrimSpace(strings.ToLower(auth[0])) != "bearer" {
			return ""
		}

		token := strings.Trim(auth[1], " ")
		token = strings.TrimSpace(token)
		return token
	} else {
		return ""
	}
}
