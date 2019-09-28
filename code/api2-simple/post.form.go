func MakeRequest() {

	formData := url.Values{
		"ID":   {"1234X"},
		"Name": {"TV LED"},
	}

	resp, err := http.PostForm("http://localhost:8181/products", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}
