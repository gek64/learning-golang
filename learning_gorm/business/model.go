package business

import "time"

type PurchasedRecord struct {
	UserId       string
	UserName     string
	ProductName  string
	SKU          string
	Price        float64
	PurchaseTime time.Time
}

type Consumption struct {
	UserId      string
	Name        string
	ProductList []string
	Total       float64
}
