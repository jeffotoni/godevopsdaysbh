package main

import (
	"fmt"
	"math"
)

func main() {
	var i interface{}

	i = "DevOps BH"
	i = 2019
	i = 9.5

	r := i.(float64)
	fmt.Println("Area do circulo: ", math.Pi*r*r)

	switch v := i.(type) {
	case int:
		fmt.Println("Int * 2=", v*2)
	case float64:
		fmt.Println("Float64/2=", v/2)
	case string:
		h := len(v) / 2
		fmt.Println("Quantidade/2 -> v[h:] + v[:h]=", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}
}
