package main

import "plugin"

func main() {
	p, err := plugin.Open("../plugin/f.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	*v.(*int) = 7
	f.(func(int))(10) // prints "Numero: 7"
}
