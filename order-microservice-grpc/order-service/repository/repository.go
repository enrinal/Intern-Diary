package repository

import (
	"context"
	pb "gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/proto/order"
	"go.mongodb.org/mongo-driver/mongo"
)


type Repository interface {
	Create(order *pb.Order) error
	GetAll() ([]*pb.Order, error)
}

type OrderRepository struct {
	Collection *mongo.Collection
}

//Create new Order
func (repo *OrderRepository) Create(order *pb.Order) error {
	_, err := repo.Collection.InsertOne(context.Background(), order)
	if err != nil {
		return err
	}
	return nil
}

//GetAll order
func (repo *OrderRepository) GetAll() ([]*pb.Order, error) {
	or, err := repo.Collection.Find(context.Background(), nil,nil)
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order
	for or.Next(context.Background()) {
		var order *pb.Order
		if err := or.Decode(&order); err != nil{
			return nil, err
		}
		orders = append(orders,order)
	}
	return orders,nil
}

