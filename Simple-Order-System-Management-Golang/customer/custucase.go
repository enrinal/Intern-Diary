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

//REGULARBUYER const for declare buyer with status = 1
const REGULARBUYER = 1

//SUBCRIPTIONBUYER const for declare buyer with status = 2
const SUBCRIPTIONBUYER = 2

func IsOrderAccept(buyer Customer) bool {
	if buyer.Status == REGULARBUYER && buyer.CountOrder <= 5 && buyer.CountItem <= 10 {
		return true
	} else if buyer.Status == SUBCRIPTIONBUYER && buyer.CountItem <= 10 {
		return true
	} else {
		return false
	}
}

func (buyer *Customer) AddOrder(numOrder int) string {
	buyer.CountOrder += numOrder
	if IsOrderAccept(*buyer) == true {
		order.NewOrder(1, "Item", order.PENDING)
		return "Order Added"
	}
	return "Order limit exceeded"
}

func (buyer *Customer) GetSumOrder() int {
	return buyer.CountOrder
}

func (buyer *Customer) AddItem(item []string) string {
	numitem := len(item)
	buyer.CountItem += numitem
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

// func (buyer *Customer) AddItem(item domain.Item) bool {
// 	if IsOrderAccept(*buyer) == true {
// 		buyer.CountItem++
// 		//domain.Order{buyer.Id,buyer.Customer,[]domain.Item{item}}
// 		return true
// 	} else {
// 		return false
// 	}
// }
