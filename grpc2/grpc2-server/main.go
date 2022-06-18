package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
	"grpc2-server/pb"
)

// Server API (Service)
type server struct{
	pb.UnimplementedAddServiceServer
}

func main() {
	// Create Listener
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	// Create gRPC Server
	grpcServer := grpc.NewServer()

	// Register Server API (Service) into gRPC Server and then the gRPC Server will has services from Server API (Service)
	pb.RegisterAddServiceServer(grpcServer, &server{})
	// reflection.Register(grpcServer)

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &pb.Response{Result: result}, nil
}
