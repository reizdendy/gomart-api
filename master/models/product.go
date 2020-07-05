package models

type Product struct {
	ProductID       string `json:"productid"`
	ProductCode     string `json:"productcode"`
	ProductName     string `json:"productname"`
	ProductCategory string `json:"productcategory"`
	ProductPrice    int    `json:"productprice"`
	Status          string `json:"status"`
}
