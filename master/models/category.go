package models

type Category struct {
	CategoryID   string `json:"categoryid"`
	CategoryName string `json:"categoryname"`
	Status       string `json:"categorystatus"`
}
