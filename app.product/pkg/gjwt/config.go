package gjwt

import "crypto/rsa"

var (

	// the keys into memory
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	// only if it will be used by files
	PathPrivate = "./private.rsa"
	PathPublic  = "./public.rsa.pub"

	ProjectTitle = "jwt Apiserver"

	ExpirationHours = 1 // Hours
	DayExpiration   = 1 // Days

)

// Structure of our server configurations
type JsonMsg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
