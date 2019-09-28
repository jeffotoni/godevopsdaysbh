package main

import (
	"encoding/json"
	"fmt"
)

type estado struct {
	Nome string  `json:"nome"`
	Area float64 `json:"area"`
}

var E []estado

func main() {
	// ---------STRUCT Unmarshal -----------------------------

	jsonEstado := []byte(`[
		{"Nome":"Acre","Area":345.66},
		{"Nome":"Amapá","Area":3445.26}
		]`)

	var ej []estado
	err := json.Unmarshal(jsonEstado, &ej)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ej)
}
