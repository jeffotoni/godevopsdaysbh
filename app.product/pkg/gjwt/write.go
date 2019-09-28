package gjwt

import (
	"encoding/json"
	"net/http"

	"github.com/jeffotoni/godevopsdasybh/app.product/pkg/util"
)

// Returns json without typing in http
func WriteJson(Status string, Msg string) {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"status":"error","message":"We could not generate the json error!"}`
		util.Print(msg)
		return
	}

	// byte
	util.Print(string(msgJson))
}

// Returns json by typing on http
func HttpWriteJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"status":"error","message":"We could not generate the json error!"}`
		w.WriteHeader(http.StatusForbidden)
		//io.WriteString(w, msg)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msgJson)
}

// Returns json by typing on http
func GetJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) string {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"status":"error","message":"We could not generate the json error!"}`
		return msg
	}
	return string(msgJson)
}
