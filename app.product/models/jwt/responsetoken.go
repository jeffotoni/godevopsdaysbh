// Go Api server
// @jeffotoni
// 2019-05-14

package models

//
// ResponseToken
//
type ResponseToken struct {

	//
	// token
	//
	Token string `json:"token"`

	Expires string `json:"expires"`

	Message string `json:"message"`
}
