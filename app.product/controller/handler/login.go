package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	mlogin "github.com/jeffotoni/godevopsdasybh/app.product/models/login"
	gjwt "github.com/jeffotoni/godevopsdasybh/app.product/pkg/gjwt"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/logrus"
	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/util"
	repol "github.com/jeffotoni/godevopsdasybh/app.product/repo/login"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodPost {
		// Do login
		SignInJson(w, r)
	} else {
		// redirect register
		jsonstr := `{"status":"error","message":"O método permitido é POST!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

// Do login JSON
func SignInJson(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != strings.ToUpper(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"status":"error","message":"O método é obrigatório! [POST]"}`
		w.Write([]byte(jsonstr))
		return
	}

	content_type := strings.ToLower(r.Header.Get("Content-Type"))
	if content_type != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		jsonstr := `{"status":"error","message":"Content-type é obrigatório!"}`
		w.Write([]byte(jsonstr))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "Login",
			"method":  "ReadAll",
		}).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"status":"error","message":"Ocorreu um erro ao ler o corpo da solicitação!"}`
		w.Write([]byte(jsonstr))
		return
	}

	var ljson mlogin.AdLoginAuth
	err = json.Unmarshal(body, &ljson)
	if err != nil {
		logrus.WithFields(logrus.Filelds{
			"version": "1.0.0",
			"host":    "goworkshop.s3wf.com",
			"handler": "Login",
			"method":  "Unmarshal",
		}).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"status":"error","message":"Ocorreu um erro ao tentar decodificar o json recebido!"}`
		w.Write([]byte(jsonstr))
		return
	}

	login := ljson.Login
	// only lower
	login = strings.TrimSpace(strings.ToLower(login))
	password := ljson.Password

	if len(login) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"status":"error","message":"O campo login é obrigatório!"}`
		w.Write([]byte(jsonstr))
		return
	}

	if len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		jsonstr := `{"status":"error","message":"O campo password é obrigatório!"}`
		w.Write([]byte(jsonstr))
		return
	}

	ok1, DLogin, code := repol.ValideGetLogin(login, password)
	// authenticates the user in the bank
	if ok1 {
		msgx := "Usuario encontrado "

		// this is required
		if DLogin.Login.Id <= 0 {
			jsonstr := `{"status":"error","message":"Usuario não encontrado!"}`
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(jsonstr))
			return
		}

		// this is required
		if len(DLogin.Login.Uuiduser) <= 0 {
			jsonstr := `{"status":"error","message":"Uuid de Usuario não encontrado!"}`
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(jsonstr))
			return
		}

		ok, tokenjwt, expires, tokenMsg := gjwt.GetTokenLogin(w, r)
		if !ok {
			jsonstr := `{"status":"error","message":"Erro ao gerar o token de acesso!"}`
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(jsonstr))
			return
		}

		// insert log acess
		repol.InsertLog(login, r.RemoteAddr)
		// struct
		DLogin.Message = msgx
		DLogin.Status = "success"

		DLogin.TokenJwt = tokenjwt
		DLogin.Expires = expires
		DLogin.TokenMsg = tokenMsg
		//DLogin.Login.TokenEnpoint = tokenEnpointsFw // token sem expirar..

		// Marshal of the structure
		b, err := json.Marshal(DLogin)
		if err != nil {
			logrus.WithFields(logrus.Filelds{
				"version": "1.0.0",
				"host":    "goworkshop.s3wf.com",
				"handler": "Login",
				"method":  "Marshal",
			}).Error(err.Error())

			jsonstr := util.Concat(`{"status":"error","message":"`, err.Error(), `"}`)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(jsonstr))
			return
		}

		// string json login
		jsonGetLogin := string(b)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonGetLogin))
		return
	}

	jsonstr := `{"status":"error","message":"Usuário não pode ser autenticado!"}`
	w.WriteHeader(code)
	w.Write([]byte(jsonstr))
	return
}
