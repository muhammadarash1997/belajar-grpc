package main

import (
	"grpc1-server/pb"
	"log"
	"net"

	"grpc1-server/chat"

	"google.golang.org/grpc"
)

func main() {
	// Create Listener
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC Server
	grpcServer := grpc.NewServer()

	// Register Server API (Service) into gRPC Server and then the gRPC Server will has services from Server API (Service)
	pb.RegisterChatServiceServer(grpcServer, &chat.Server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
