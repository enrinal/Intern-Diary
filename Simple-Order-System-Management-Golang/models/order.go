package models

type Order struct {
	ID       int
	Item     string
	Customer Customer
	Status   int
}
