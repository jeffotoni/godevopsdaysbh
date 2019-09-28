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

	// token 64
	authToken64 := strings.SplitN(tokenBase64, ":", 2)
	if len(authToken64) != 2 || authToken64[0] == "" || authToken64[1] == "" {
		msgjson = GetJson(w, "error", "Invalido key!", http.StatusUnauthorized)
		return
	}

	tokenUserEnc := authToken64[0]
	keyUserEnc := authToken64[1]

	// User, Login byte
	tokenUserDecode, _ := b64.StdEncoding.DecodeString(tokenUserEnc)

	// key user byte
	keyUserDec, _ := b64.StdEncoding.DecodeString(keyUserEnc)

	// User, Login string
	tokenUserDecodeS = strings.TrimSpace(strings.Trim(string(tokenUserDecode), " "))

	// key user, string
	keyUserDecS := strings.TrimSpace(strings.Trim(string(keyUserDec), " "))

	UserR, _ := b64.StdEncoding.DecodeString(cf.X_KEY_1)
	PassR, _ := b64.StdEncoding.DecodeString(cf.X_KEY_2)

	// Validate user and password in the database
	if tokenUserDecodeS == string(UserR) && keyUserDecS == string(PassR) {
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
