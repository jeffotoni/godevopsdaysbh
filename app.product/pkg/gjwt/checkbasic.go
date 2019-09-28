package gjwt

import (
	b64 "encoding/base64"
	cf "github.com/jeffotoni/godevopsdasybh/app.product/config"
	"net/http"
	"strings"
)

// validates and generates jwt token
func CheckBasic(w http.ResponseWriter, r *http.Request) (ok bool, msgjson, tokenUserDecodeS string) {

	ok = false
	msgjson = `{"status":"error","message":"tentando autenticar usuário!"}`

	// Authorization Basic base64 Encode
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		msgjson = GetJson(w, "error", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	tokenBase64 := strings.Trim(auth[1], " ")
	tokenBase64 = strings.TrimSpace(tokenBase64)

	// User, Login byte
	tokenUserDecode, _ := b64.StdEncoding.DecodeString(tokenBase64)
	tokenUserDecodeS = string(tokenUserDecode)

	tokenLocal := cf.X_KEY_1 + ":" + cf.X_KEY_2

	// Validate user and password in the database
	if tokenUserDecodeS == tokenLocal {
		ok = true
		return ok, `{"status":"ok"}`, tokenUserDecodeS
	} else {
		stringErr := "Usuário e chaves inválidas"
		//+ auth[0] + " - " + auth[1]
		msgjson = GetJson(w, "error", stringErr, http.StatusUnauthorized)
		return ok, msgjson, tokenUserDecodeS
	}

	defer r.Body.Close()
	return ok, msgjson, tokenUserDecodeS
}
