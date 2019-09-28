// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repol

import (
	"strings"

	mlogin "github.com/jeffotoni/godevopsdasybh/app.product/models/login"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/crypt"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/pg"
)

func ValideGetLogin(email, password string) (bool, *mlogin.GetLoginData, int) {

	Db := pg.ConnectNew()

	// garantir caixa baixa
	email = strings.ToLower(email)
	row := Db.QueryRow(`SELECT logi_senha,logi_id,logi_uuid,logi_nome,logi_last_name,
		logi_avatar_dominio,logi_data_criacao::text,logi_avatar_type 
		FROM public.user WHERE lower(logi_email)=$1`, email)

	var logi_id int64
	var logi_uuid string // required
	var logi_nome string
	var logi_senha string
	var logi_data_criacao, logi_avatar_type *string
	var logi_avatar_dominio, logi_last_name *string

	errqy := row.Scan(&logi_senha, &logi_id, &logi_uuid, &logi_nome, &logi_last_name, &logi_avatar_dominio, &logi_data_criacao, &logi_avatar_type)

	gl := mlogin.GetLoginData{}

	if errqy != nil {
		// logrus.WithFields(logrus.Filelds{
		// 	"version": "1.0.0",
		// 	"host":    "goworkshop.s3wf.com",
		// 	"handler": "ValideGetLogin",
		// 	"method":  "row.Scan",
		// }).Error(errqy.Error())
		return false, nil, 500
	}

	if len(logi_senha) > 0 {
		if !crypt.CheckBlowfish(password, logi_senha) {
			return false, nil, 500
		}
	}

	gl.Login.Id = logi_id
	gl.Login.Uuiduser = logi_uuid
	gl.Login.Name = logi_nome

	if logi_avatar_dominio != nil {
		gl.Login.AvatarUrl = *logi_avatar_dominio
	}

	if logi_last_name != nil {
		gl.Login.LastName = *logi_last_name
	}

	if logi_data_criacao != nil {
		gl.Login.DataStart = *logi_data_criacao
	}
	if logi_avatar_type != nil {
		gl.Login.AvatarType = *logi_avatar_type
	}

	return true, &gl, 200
}
