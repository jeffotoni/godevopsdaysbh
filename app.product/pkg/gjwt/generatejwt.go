package gjwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	cf "github.com/jeffotoni/godevopsdasybh/app.product/config"
	"github.com/jeffotoni/godevopsdasybh/app.product/models/jwt"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
	"strconv"
	"time"
)

// jwt generateJwt
func generateJwt(model models.User) (string, string) {

	// convert int to duration
	HourxMonth := time.Duration(ExpirationHours * DayExpiration)

	// Generating date validation to return to the user
	Expires := time.Now().Add(time.Hour * HourxMonth).Unix()

	// convert int6
	ExpiresInt64, _ := strconv.ParseInt(fmt.Sprintf("%v", Expires), 10, 64)

	// convert time unix to Date RFC
	ExpiresDateAll := time.Unix(ExpiresInt64, 0)

	// Date
	ExpiresDate := ExpiresDateAll.Format(cf.LayoutDate)

	// claims Token data, the header
	claims := models.Claim{
		User: model.Login,
		StandardClaims: jwt.StandardClaims{
			// Expires in 24 hours * 10 days
			ExpiresAt: Expires,
			Issuer:    ProjectTitle,
		},
	}

	// Generating token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Transforming into string
	tokenString, err := token.SignedString(PrivateKey)
	if err != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "generateJwt",
			"method":  "token.SignedString",
		}).Error(err.Error())
		return "Could not sign the token!", "2006-01-02"
	}

	// return token string
	return tokenString, ExpiresDate
}
