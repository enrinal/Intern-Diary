package main

import (
	"order/customer"
)

func main() {
	buyer := customer.NewCustomer(1, "Enrinal", 1, 1, 1)
	buyer.AddOrder(1)
}
