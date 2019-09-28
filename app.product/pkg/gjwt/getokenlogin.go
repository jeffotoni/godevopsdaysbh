package gjwt

import (
	"github.com/jeffotoni/godevopsdasybh/app.product/models/jwt"
	"net/http"
)

// validates and generates jwt token
func GetTokenLogin(w http.ResponseWriter, r *http.Request) (ok bool, token, expires, tokenMsg string) {

	ok = false
	defer r.Body.Close()
	ok, _, tokenUserDecodeS := CheckBasic(w, r)

	if ok {

		var model models.User
		model.Login = tokenUserDecodeS
		model.Password = ""
		model.Role = "admin"
		tokenMsg = "use o token para acessar os endpoints!"
		token, expires = generateJwt(model)

		ok = true
		return

	} else {
		//msgjson = GetJson(w, "error", msgjson, http.StatusUnauthorized)
		return
	}
}
