package main

import "plugin"
import "io"
import "strings"
import "os"

func main() {
	p, err := plugin.Open("../plugin/randseed.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("RandSeed")
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, strings.NewReader(f.(func() string)()))
}
