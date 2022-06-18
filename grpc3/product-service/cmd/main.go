package main

import (
	"fmt"
	"log"
	"net"
	"product-service/pkg/config"
	"product-service/pkg/db"
	"product-service/pkg/pb"
	"product-service/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	handler := db.Init(c.DBUrl)

	// Create Listener
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}
	fmt.Println("Product Svc on", c.Port)

	// Create Server API
	s := services.Server{
		Handler: handler,
	}

	// Create gRPC Server which still has no services
	grpcServer := grpc.NewServer()

	// Register Server API (Service) into gRPC Server and the gRPC Server will has services from Server API (Service)
	pb.RegisterProductServiceServer(grpcServer, &s)

	// Serve accepts incoming connections on the Listener
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
