// Go Api server
// @jeffotoni

package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim structure, where we will use
// to validate our token with jwt

//
// jwt
//
type Claim struct {

	//
	//
	//
	User string `json:"user"`

	//
	//
	//
	jwt.StandardClaims
}
