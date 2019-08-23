package main

import (
	"encoding/json"
	micro "github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "gitlab.warungpintar.co/enrinal/hermes-son-of-zeus/order-service/proto/order"
)

const (
	address         = "localhost:50051"
	defaultFilename = "order.json"
)

func parseFile(file string) (*pb.Order, error) {
	var order *pb.Order
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &order)
	return order, err
}

func main() {

	service := micro.NewService(micro.Name("hermes.cli.order"))
	service.Init()

	client := pb.NewOrderServiceClient("hermes.service.order", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateOrder(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetOrder(context.Background(), &pb.GetOrderReq{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _,v := range getAll.Orders {
		log.Println(v)
	}
}
