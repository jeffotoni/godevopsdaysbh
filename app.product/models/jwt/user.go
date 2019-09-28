// Go Api server
// @jeffotoni
// 2019-05-14

package models

//
// User structure
//
type User struct {

	//
	//
	//
	Login string `json:"login"`

	//
	//
	//
	Password string `json:"password,omitempty"`

	//
	//
	//
	Codemp string `json:"codemp"`

	//
	//
	//
	Role string `json:"role"`
}
