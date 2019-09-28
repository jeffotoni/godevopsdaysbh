package gjwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeffotoni/godevopsdasybh/app.product/cert"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
)

// jwt init
func init() {

	var errx error
	privateByte := []byte(cert.RSA_PRIVATE)
	PrivateKey, errx = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if errx != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "init",
			"method":  "jwt.ParseRSAPrivateKeyFromPEM",
		}).Error(errx.Error())
		return
	}

	publicByte := []byte(cert.RSA_PUBLIC)
	PublicKey, errx = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if errx != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "init",
			"method":  "jwt.ParseRSAPublicKeyFromPEM",
		}).Error(errx.Error())
		return
	}
}
