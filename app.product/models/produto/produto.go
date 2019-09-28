package mproducts

type Categories struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

type Associations struct {
    Platform string `json:"platform"`
    Status   string `json:"status"`
}

type Specifications struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type Product struct {
    Sku               string   `json:"sku"`
    Name              string   `json:"name"`
    Description       string   `json:"description"`
    Status            string   `json:"status"`
    Removed           bool     `json:"removed"`
    Qty               int64    `json:"qty"`
    Price             float64  `json:"price"`
    Promotional_price float64  `json:"promotional_price"`
    Cost              *float64 `json:"cost"`
    Weight            float64  `json:"weight"`
    Height            float64  `json:"height"`
    Width             float64  `json:"width"`
    Length            float64  `json:"length"`
    Brand             string   `json:"brand"`
    Ean               string   `json:"ean"`
    Nbm               *string  `json:"nbm"`

    Categories     []Categories     `json:"categories"`
    Images         []string         `json:"images"`
    Specifications []Specifications `json:"specifications"`
    Associations   []Associations   `json:"associations"`
}

//////// SOMENTE PARA POSTS
type Post2Product struct {
    Product Product `json:"product"`
}

type GetProduct struct {
    Product Product `json:"products"`
}

type GetProduct2 struct {
    P     []Product `json:"products"`
    Total int       `json:"total"`
}
