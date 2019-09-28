package crypt

import (
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Blowfish(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "Blowfish",
			"method":  "bcrypt.GenerateFromPassword",
		}).Error(err.Error())
		//panic(err)
	}
	return string(bytes)
}

func CheckBlowfish(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
