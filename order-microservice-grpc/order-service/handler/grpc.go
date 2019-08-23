package handler

import (
	"context"
	pb "gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/proto/order"
	"gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/repository"
)

type OrderHandler struct {
	repository.Repository
}

//CreateOrder
func (s *OrderHandler) CreateOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {
	err := s.Repository.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Order = req
	return nil
}

//GetOrder
func (s *OrderHandler) GetOrder(ctx context.Context, req *pb.GetOrderReq, res *pb.Response) error {
	orders, err := s.Repository.GetAll()
	if err != nil {
		return err
	}
	res.Orders = orders
	return nil
}
