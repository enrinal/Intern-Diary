package usecase

import (
	"fmt"
	"order/domain"
)

type User struct {
	Id             int
	IsSeller       bool
	CountPembelian int
	CountItem      int
	Status         int
	Customer       domain.Customer
}
type Order domain.Order

func IsPembelian(pembeli User) bool {
	if pembeli.Status == 1 && pembeli.CountPembelian <= 5 && pembeli.CountItem <= 10 {
		return true
	} else if pembeli.Status == 2 && pembeli.CountItem <= 10 {
		return true
	} else {
		return false
	}
}

func (pembeli *User) BeliBarang() bool {
	if IsPembelian(*pembeli) == true {
		pembeli.CountPembelian++
		return true
	} else {
		return false
	}
}

func (pembeli *User) AmbilBarang(item domain.Item) bool {
	if IsPembelian(*pembeli) == true {
		pembeli.CountItem++
		//domain.Order{pembeli.Id,pembeli.Customer,[]domain.Item{item}}
		return true
	} else {
		return false
	}
}

func PrintUser(pembeli User) {
	fmt.Printf("%d %s %d %d %d \n", pembeli.Id, pembeli.Customer.Name, pembeli.CountItem, pembeli.CountPembelian, pembeli.Status)
}

func (order *Order) ProsesOrder() bool {
	order.Status = "process"
	return true
}

func (order *Order) DikirimOrder() bool {
	order.Status = "sent"
	return true
}

func (order *Order) DiterimaOrder() bool {
	order.Status = "delivered"
	return true
}

func PrintOrder(order Order) {
	fmt.Println("Order id: ", order.Id)
	fmt.Println("Customer name: ", order.Customer.Name)
	fmt.Println("Status : ", order.Status)
}
