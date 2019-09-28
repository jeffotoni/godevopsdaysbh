package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)
	// 3

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi)
	// 3.1412

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str)

}
