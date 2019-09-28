package gjwt

import (
	"net/http"
)

//check if basic token exists
func CheckJwt(w http.ResponseWriter, r *http.Request) bool {
	if !tokenJwtClaimsValid(w, r) {
		return false
	}
	return true
}
