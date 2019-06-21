package models

type Cart struct {
	IDCart int64
	IDCust int64
	Items  []Item
	Qty    int64
	Prize  int64
}
