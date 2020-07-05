package models

type Report struct {
	TransactionDate string `json:"transactiondate"`
	ProductName     string `json:"productname"`
	ProductPrice    int    `json:"productprice"`
	Quantity        int    `json:"quantity"`
	Total           int    `json:"total"`
}
