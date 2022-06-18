package chat

import (
	"context"
	"log"
	"grpc1-server/pb"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (this *Server) SayHello(ctx context.Context, message *pb.RequestMessage) (*pb.ResponseMessage, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &pb.ResponseMessage{Body: message.Body}, nil
}
