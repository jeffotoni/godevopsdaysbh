// Go Api server
// @jeffotoni

package mlogin

type Login struct {
	Login        string `json:"login,omitempty"`
	Id           int64  `json:"id,omitempty"`
	Uuiduser     string `json:"uuiduser,omitempty"`
	Email        string `json:"email,omitempty"`
	AvatarUrl    string `json:"avatarurl,omitempty"`
	AvatarExt    string `json:"avatarext,omitempty"`
	AvatarType   string `json:"avatartype,omitempty"`
	Level        int64  `json:"level,omitempty"`
	Name         string `json:"name,omitempty"`
	Token        string `json:"token,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	DataStart    string `json:"datastart,omitempty"`
	TokenEnpoint string `json:"tokenenpoint"`
	Codemp       int64  `json:"codemp"`
}

type GetLoginData struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	TokenJwt string `json:"tokenjwt"`
	Expires  string `json:"expires"`
	TokenMsg string `json:"tokenmsg"`
	Login    Login
}
