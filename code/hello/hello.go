package main

import "io"
import "os"
import "strings"

func main() {
	io.Copy(os.Stdout, strings.NewReader("Hello World!"))
}
