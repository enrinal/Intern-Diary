package models

type Cart struct {
	Items map[string]Item
}

type Item struct {
	Id    string
	Name  string
	Price string
	Qty   int
}