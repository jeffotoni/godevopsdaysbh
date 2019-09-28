
func MakeNewRequestFull() {

	bodyPost := `{"msg":"texto em json aqui..."}`

	var apiUrl string = "http://localhost:8181"

	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	defer afterFuncTimer.Stop()

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = "/products"
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(bodyPost))
	req = req.WithContext(ctx)

	Authorization := crypt.Base64Enc(util.Concat(usuario, ":", password))
	req.Header.Add("Authorization", util.Concat("Basic ", Authorization))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 &&
		resp.StatusCode <= 299 {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

/**
FORMULARIO
data := url.Values{}
data.Add("grant_type", "authorization_code")
data.Add("client_id", appid)
data.Add("client_secret", secret)
data.Add("code", UserCode)
data.Add("redirect_uri", REDIRECT_URL)
req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // This makes it work
**/

// objProdSend := mprods.ProdSendPai{
// 		Product_Send: mprods.ProductSend{
// 			Sku:         codprod,
// 			Marketplace: marketplace,
// 			Codident:    codident,
// 		},
// 	}

// 	bdata, err := json.Marshal(objProdSend)
// 	if err != nil {
// 		log.Println("Marshal: ", err)
// 		return http.StatusBadRequest, err
// 	}

// 	u, _ := url.ParseRequestURI(apiUrl)
// 	u.Path = resource
// 	urlStr := u.String()
// 	req, err := http.NewRequest("PUT", urlStr, bytes.NewBuffer(bdata))
// 	req = req.WithContext(ctx)

// 	req.Header.Add("X-User-Email", usuario)
// 	req.Header.Add("X-Api-key", password)
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Content-Length", strconv.Itoa(len(string(bdata))))
//req.Header.Add("Content-Length", strconv.Itoa(len(string(bdata.Bytes()))))

/////// outra forma de envio
//bdata := new(bytes.Buffer)
//json.NewEncoder(bdata).Encode(objProdSend)
// req, err := http.NewRequest("PUT", urlStr, bdata)
//req.Header.Add("Content-Length", strconv.Itoa(len(string(bdata.Bytes()))))
////////////////