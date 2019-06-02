package usecase

import (
	"fmt"
	"order/domain"
)

type User struct {
	Id         int
	IsSeller   bool
	CountOrder int
	CountItem  int
	Status     int
	Customer   domain.Customer
}
type Order domain.Order

func IsOrderAccept(buyer User) bool {
	if buyer.Status == 1 && buyer.CountOrder <= 5 && buyer.CountItem <= 10 {
		return true
	} else if buyer.Status == 2 && buyer.CountItem <= 10 {
		return true
	} else {
		return false
	}
}

func (buyer *User) Order() bool {
	if IsOrderAccept(*buyer) == true {
		buyer.CountOrder++
		return true
	} else {
		return false
	}
}

func (buyer *User) AddItem(item domain.Item) bool {
	if IsOrderAccept(*buyer) == true {
		buyer.CountItem++
		//domain.Order{buyer.Id,buyer.Customer,[]domain.Item{item}}
		return true
	} else {
		return false
	}
}

func PrintUser(buyer User) {
	fmt.Printf("%d %s %d %d %d \n", buyer.Id, buyer.Customer.Name, buyer.CountItem, buyer.CountOrder, buyer.Status)
}

func (order *Order) OrderProcess() bool {
	order.Status = 2
	return true
}

func (order *Order) OrderSend() bool {
	order.Status = 3
	return true
}

func (order *Order) OrderDelivered() bool {
	order.Status = 4
	return true
}

func PrintOrder(order Order) {
	fmt.Println("Order id: ", order.Id)
	fmt.Println("Customer name: ", order.Customer.Name)
	fmt.Println("Status : ", order.Status)
}
