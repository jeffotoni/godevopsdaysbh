package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	////////// Decode
	// "Parents":["Arthur","Francisco"]
	jsonData := `[{"Name":"Jef","Age":36},{"Name":"Go","Age":10}]`
	var m []map[string]interface{}

	// uma linha decoder
	json.NewDecoder(strings.NewReader(jsonData)).Decode(&m)
	fmt.Println(&m)
}
