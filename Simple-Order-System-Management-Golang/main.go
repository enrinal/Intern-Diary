package main

import (
	"order/domain"
	"order/usecase"
)

func main() {
	cust := domain.Customer{1, "Enrinal"}
	user := usecase.User{1, false, 1, 1, usecase.REGULARBUYER, cust}
	sapi := domain.Item{1, "Sapi", true}
	rumah := domain.Item{3, "Rumah", true}
	mobil := domain.Item{2, "Mobil", true}
	order := usecase.Order{1, user.Customer, []domain.Item{sapi, mobil, rumah}, 1}
	user.AddItem(sapi)
	usecase.PrintUser(user)
	order.OrderProcess()
	usecase.PrintOrder(order)
}
