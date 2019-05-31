package main

import (
	"order/domain"
	"order/usecase"
)

func main() {
	cust := domain.Customer{1, "Enrinal"}
	user := usecase.User{1, false, 1, 1, 1, cust}
	sapi := domain.Item{1, "Sapi", true}
	mobil := domain.Item{2, "Mobil", true}
	order := usecase.Order{1, user.Customer, []domain.Item{sapi, mobil}, "pending"}
	user.AmbilBarang(sapi)
	usecase.PrintUser(user)
	order.ProsesOrder()
	usecase.PrintOrder(order)
}
