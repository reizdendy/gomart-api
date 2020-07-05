package models

type Transaction struct {
	TransactionId        string `json:"transactionid"`
	TransactionDate      string `json:"transactiondate"`
	TransactionProductId string `json:"productid"`
	Quantity             int    `json:"qty"`
}
