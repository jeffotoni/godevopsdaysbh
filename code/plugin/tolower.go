// Go in action
// @jeffotoni
// 2019-08-10
package main

import (
	"strings"
)

type tolower string

func (t tolower) Tolower(name string) string {
	return strings.ToLower(name)
}

// this is exported
var Tolower tolower
