package main

import (
	"fmt"
	"log"
	"net"

	"order-service/pkg/client"
	"order-service/pkg/config"
	"order-service/pkg/db"
	"order-service/pkg/pb"
	"order-service/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	handler := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	fmt.Println("Order Svc on", c.Port)

	// Create Server API
	s := service.Server{
		Handler:          handler,
		ProductSvc: productSvc,
	}

	// Create gRPC Server which still has no services
	grpcServer := grpc.NewServer()

	// Register Server API (Service) into gRPC Server and the gRPC Server will has services from Server API (Service)
	pb.RegisterOrderServiceServer(grpcServer, &s)

	// Serve accepts incoming connections on the Listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
