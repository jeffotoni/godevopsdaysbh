// Go Api server
// @jeffotoni

package util

import (
	"net/http"
	"regexp"
	"strings"
)

// regex email valid
func RegexEmail(nameregex string) bool {
	// regex for my endpoint
	var loginRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Contains email valid
	if loginRegex.MatchString(nameregex) {
		return true
	} else {
		return false
	}
}

// generating endpoint and endpoint nameregex
func EndpointRegex(r *http.Request) (string, string) {
	vetPath := strings.Split(r.URL.Path, "/")
	lastpos := len(vetPath) - 1
	nameregex := vetPath[lastpos]
	endpoint := ""
	for _, v := range vetPath {
		if v != nameregex {
			endpoint += "/" + v
		}
	}
	// endpoint
	endpoint = "/" + strings.Trim(endpoint, "/")
	return endpoint, nameregex
}

// generating endpoint and endpoint nameregex
func EndpointRegexTwo(r *http.Request) (string, string) {
	vetPath := strings.Split(r.URL.Path, "/")
	lastpos := len(vetPath) - 2
	nameregex := vetPath[lastpos]
	endpoint := ""
	for _, v := range vetPath {
		if v != nameregex {
			endpoint += "/" + v
		}
	}
	// endpoint
	endpoint = "/" + strings.Trim(endpoint, "/")
	return endpoint, nameregex
}
