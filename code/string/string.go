package main

import "strings"

func Native() string {
	var s string
	for i := 0; i < 100; i++ {
		//s = s + `go`
		s += string(i)
	}

	return s
}

func Builder() string {
	b := strings.Builder{}
	//b.Grow(60)
	for i := 0; i < 100; i++ {
		b.WriteString(string(i))
	}
	return b.String()

}

func main() {}
