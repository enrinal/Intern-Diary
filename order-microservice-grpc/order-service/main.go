package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/datastore"
	pb "gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/proto/order"
	"gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/repository"
	or "gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/handler"
	"log"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("hermes.service.order"),
		micro.Address("127.0.0.1:54782"),
	)

	srv.Init()

	uri := defaultHost
	client, err := datastore.CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	orderCollection := client.Database("hermes").Collection("orders")

	repo := &repository.OrderRepository{orderCollection}
	h := &or.OrderHandler{repo}


	//Register handler
	pb.RegisterOrderServiceHandler(srv.Server(), h)

	//Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
