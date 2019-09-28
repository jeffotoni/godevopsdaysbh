package main

import (
	"encoding/json"
	"fmt"
)

type jsoninput []struct {
	Name string `json:"name"`
}

func main() {

	// json in memory
	// this is an array of values
	resp := `[{"name":"Jeffotoni"},{"name":"Pike"}]`

	// initialization struct
	data := &jsoninput{}

	// Unmarshal in bytes
	_ = json.Unmarshal([]byte(resp), data)

	// loop to see the values in the fields
	// loop in struct
	for _, value := range *data {
		fmt.Println(value.Name)
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
