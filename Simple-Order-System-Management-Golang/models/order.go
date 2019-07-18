package models

type Order struct {
	ID     int64  `json:"id"`
	IDCust int64  `json:"idcust"`
	Item   string `json:"item"`
	Status int64  `json:"status"`
}
