// Go Api server
// @jeffotoni
// 2019-05-14

package mlogin

type AdLoginAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	//Keep_session bool   `json:keep_session`
}
