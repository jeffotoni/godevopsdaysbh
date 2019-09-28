package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Bird struct {
		Species     string `json:"birdType"`
		Description string `json:"what it does"`
	}

	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks}

}
