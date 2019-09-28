package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// DINAMIC
	jsonDataByte := []byte(`{"Name":"Rob Pike","Age":46,"Parents":["GO","Occam"]}`)
	var vv interface{}
	json.Unmarshal(jsonDataByte, &vv)
	datamap := vv.(map[string]interface{})

	for k, v := range datamap {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}
