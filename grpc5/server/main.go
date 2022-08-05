package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/muhammadarash1997/grpc5/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedWearableServiceServer
}

func newServer() *Server {
	return &Server{}
}

func (s *Server) BeatsPerSecond(req *pb.BeatsPerSecondRequest, stream pb.WearableService_BeatsPerSecondServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return status.Errorf(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)
			value := rand.Intn(80)
			err := stream.SendMsg(&pb.BeatsPerSecondResponse{
				Value:  uint32(value),
				Second: uint32(time.Now().Second()),
			})
			if err != nil {
				return status.Errorf(codes.Canceled, "Stream has ended")
			}
		}
	}
}

func main() {
	// Create Listener
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create gRPC Server
	grpcServer := grpc.NewServer()

	// Register Service Server into gRPC Server and then the gRPC Server will has services
	pb.RegisterWearableServiceServer(grpcServer, newServer())

	// Listen and Serve of Listener and gRPC Server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Error serve grpc", err.Error())
	}
}
