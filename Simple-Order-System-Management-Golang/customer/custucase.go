package customer

import (
	"fmt"
	"order/order"
)

func NewCustomer(ID int, Name string, CountOrder int, CountItem int, Status int) *Customer {
	cust := new(Customer)
	cust.ID = ID
	cust.Name = Name
	cust.CountOrder = CountOrder
	cust.CountItem = CountItem
	cust.Status = Status
	return cust
}

const (
	RegularBuyer = iota + 1
	SubcriptionBuyer
)

func IsOrderAccept(buyer Customer) bool {
	if (buyer.Status == RegularBuyer && buyer.CountOrder <= 5 && buyer.CountItem <= 10) ||
		(buyer.Status == SubcriptionBuyer && buyer.CountItem <= 10) {
		return true
	}
	return false
}

func (buyer *Customer) AddOrder(numOrder int) string {
	buyer.CountOrder += numOrder
	if IsOrderAccept(*buyer) == true {
		order.NewOrder(1, "Item", order.Pending)
		return "Order Added"
	}
	return "Order limit exceeded"
}

func (buyer *Customer) GetSumOrder() int {
	return buyer.CountOrder
}

func (buyer *Customer) AddItem(item []string) string {
	buyer.CountItem += len(item)
	if IsOrderAccept(*buyer) == true {
		return "Item Added"
	}
	return "Your basket limit exceeded"
}

func (buyer *Customer) GetSumItem() int {
	return buyer.CountItem
}

func PrintCust(buyer Customer) {
	fmt.Printf("%d %s %d %d %d \n", buyer.ID, buyer.Name, buyer.CountItem, buyer.CountOrder, buyer.Status)
}

func PrintOrder(order order.Order) {
	fmt.Println("Order id: ", order.ID)
	fmt.Println("Status : ", order.Status)
}
