package gjwt

import (
	"encoding/json"
	"github.com/jeffotoni/godevopsdasybh/app.product/models/jwt"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
	"net/http"
)

// validates and generates jwt token
func GetToken(w http.ResponseWriter, r *http.Request) (ok bool, msgjson string) {

	ok = false
	defer r.Body.Close()
	ok, msgjson, tokenUserDecodeS := CheckBasic(w, r)

	if ok {
		var model models.User
		model.Login = tokenUserDecodeS
		model.Password = ""
		model.Role = "admin"
		tokenMsg := "Use o token para acessar os endpoints!"
		token, expires := generateJwt(model)
		result := models.ResponseToken{token, expires, tokenMsg}
		jsonResult, err := json.Marshal(result)

		if err != nil {
			logrus.WithFields(logrus.Filelds{
				"version": "1.0.0",
				"host":    "goworkshop.s3wf.com",
				"handler": "GetToken",
				"method":  "json.Marshal",
			}).Error(err.Error())
			msgjson = GetJson(w, "error", "Erro ao gerar Json!", http.StatusUnauthorized)
			return
		}
		ok = true
		return ok, string(jsonResult)

		// Output
		//
		/**
		{
		  "Token": "39a3099b45634f6eb511991fddde83752_v2",
		  "Expires": "2026-09-14",
		  "Message": "use the token to access the endpoints"
		}
		*/
	} else {
		msgjson = GetJson(w, "error", msgjson, http.StatusUnauthorized)
		return ok, msgjson
	}
}
