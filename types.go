package main

type PriceResponse struct {
	Price  float64
	Ticker string
}

type Product struct {
	ProductId        int    `json:"productId"`
	ProductName      string `json:"productName"`
	ProductBuyPrice  int64  `json:"productbuyPrice"`
	ProductSellPrice int64  `json:"productSellPrice"`
}
