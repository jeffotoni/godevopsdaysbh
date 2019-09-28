// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repol

import (
	"strings"

	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
	pg "github.com/jeffotoni/godevopsdasybh/app.product/pkg/psql"
)

func InsertLog(login, ip string) bool {

	Db := pg.ConnectNew()

	// garantir caixa baixa
	login = strings.ToLower(login)

	//data := time.Now().Format(cf.LayoutDate)
	//hora := time.Now().Format(cf.LayoutHour)

	acess_ip := ip
	acess_remoto := ip
	acess_user := login
	acess_user_up := login
	acess_count := 0
	acess_client_msg := "inserindo logs de acesso via client"
	acess_client := 1

	insert := `INSERT INTO acessolog(acess_ip,acess_user,acess_user_up,acess_count,acess_remoto,acess_client_msg,acess_client) 
	VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := Db.Exec(insert, acess_ip, acess_user, acess_user_up, acess_count, acess_remoto, acess_client_msg, acess_client)

	if err != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "InsertLog",
			"method":  "Db.Exec",
		}).Error(err.Error())
		return false
	}

	return true
}
