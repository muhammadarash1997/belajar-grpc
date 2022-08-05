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

	// Register Service Server into gRPC Server so that the gRPC Server will has services
	pb.RegisterChatServiceServer(grpcServer, &chat.Server{})
	log.Printf("server listening at %v", lis.Addr())

	// Listen and Serve of Listener and gRPC Server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
