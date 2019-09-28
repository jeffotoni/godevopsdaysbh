func MakeRequest() {

	client := http.Client{}
	request, err := http.NewRequest("GET", "https://localhost:8181/products", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}

func MakeRequest2() {

	client := http.Client{}
	resp, err := client.Get("https://localhost:8181/products")
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}
