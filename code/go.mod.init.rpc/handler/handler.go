package handler

import (
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	// Do auth here
}

func UserGetAll(w http.ResponseWriter, r *http.Request) {

	// Do UserGetAll here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"ok"}`))
}

func UserGet(w http.ResponseWriter, r *http.Request) {

	// Do UserGet here
}

func UserDel(w http.ResponseWriter, r *http.Request) {

	// Do UserDel here
}

func UserPut(w http.ResponseWriter, r *http.Request) {

	// Do UserPut here
}
