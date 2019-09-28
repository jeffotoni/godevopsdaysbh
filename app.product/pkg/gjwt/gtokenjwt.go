package gjwt

import (
	b64 "encoding/base64"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	cf "github.com/jeffotoni/godevopsdasybh/app.product/config"
	"github.com/jeffotoni/godevopsdasybh/app.product/models/jwt"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
)

// GtokenJwt
// here it will check the header if there are the keys
// to authenticate and it will check with the system and generate
// the token to access all the endpoints
func GtokenJwt(w http.ResponseWriter, r *http.Request) bool {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		HttpWriteJson(w, "error", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return false
	}

	token := strings.Trim(auth[1], " ")
	strings.TrimSpace(token)

	// star
	parsedToken, err := jwt.ParseWithClaims(token, &models.Claim{}, func(*jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})

	if err != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "GtokenJwt",
			"method":  "ParseWithClaims",
		}).Error(err.Error())
	}

	if err != nil || !parsedToken.Valid {
		HttpWriteJson(w, "error", "Seu token expirou!", http.StatusUnauthorized)
		return false
	}

	UserR, _ := b64.StdEncoding.DecodeString(cf.X_KEY_1)

	claims, ok := parsedToken.Claims.(*models.Claim)
	if !ok || claims.User != string(UserR) {
		msgjson := `{"status":"error","message":"Existe alguma coisa estranha em seu token!"}`
		//io.WriteString(w, msgjson)
		w.Write([]byte(msgjson))
		return false
	}

	return true
}
