package models

type Order struct {
	Orderid             int  `json:"orderId"`
	ProductId           int  `json:"productId"`
	Amount              int  `json:"amount"`
	IsProfitDistributed bool `json:"isProfitDistributed"`
}

type OrderWithProduct struct {
	Orderid             int     `json:"orderId"`
	ProductId           Product `json:"productId"`
	Amount              int     `json:"amount"`
	IsProfitDistributed bool    `json:"isProfitDistributed"`
}
