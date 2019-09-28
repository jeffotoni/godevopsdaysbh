package main

import (
	"io"
	"os"
	"strings"
)

func Println(text string) {
	io.Copy(os.Stdout, strings.NewReader(text))
}

func main() {
	Println("\nOla Workshop2.0 ...")
}
