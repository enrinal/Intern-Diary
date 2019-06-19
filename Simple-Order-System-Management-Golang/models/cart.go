package models

type Cart struct {
	Customer Customer
	Items    map[string]Item
}
