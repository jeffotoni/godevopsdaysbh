package main

import "fmt"

type Produto struct {
	Id    int64
	Nome  string
	Preco float64
}

func main() {

	type MyString string

	const nome string = "@jeffotoni"

	var p Produto
	var vars MyString
	var t string
	var a int

	vars = MyString(nome)
	t = "ola"
	a = 100
	p.Id = 1000
	p.Nome = "TV LG"
	p.Preco = 2500.00

	fmt.Println(t)
	fmt.Println(a)
	fmt.Println(nome)
	fmt.Println(vars)
	fmt.Println(p)
}
