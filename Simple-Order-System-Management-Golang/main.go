package main

import (
	"order/domain"
	"order/usecase"
)

const REGULARBUYER = 1
const SUBCRIPTIONBUYER = 2
const PENDING = 1
const PROCESS = 2
const DDELIVERED = 3
const SEND = 2

func main() {
	cust := domain.Customer{1, "Enrinal"}
	user := usecase.User{1, false, 1, 1, REGULARBUYER, cust}
	sapi := domain.Item{1, "Sapi", true}
	mobil := domain.Item{2, "Mobil", true}
	order := usecase.Order{1, user.Customer, []domain.Item{sapi, mobil}, 1}
	user.AddItem(sapi)
	usecase.PrintUser(user)
	order.OrderProcess()
	usecase.PrintOrder(order)
}
