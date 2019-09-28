package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonByte := []byte(`[{"Name":"Occan","Year":1983},{"Name":"Go","Year":2009}]`)
	var mm []map[string]interface{}

	//////// outra forma de fazer decode.. Unmarshal
	if err := json.Unmarshal(jsonByte, &mm); err != nil {
		fmt.Println("json error: ", err)
		return
	}

	fmt.Println(&mm)
	fmt.Println(mm[0]["Name"])
	fmt.Println(mm[0]["Year"])
}
