package main

import (
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/muhammadarash1997/my-chatting/pb"

	"google.golang.org/grpc"
)

type message struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageHandler struct {
	sync.Mutex
	MessageQue []message
}

var messageHandlerObject = messageHandler{}

type server struct {
	pb.UnimplementedMyChattingServer
}

func (this *server) ChatService(stream pb.MyChatting_ChatServiceServer) error {
	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	go this.receiveFromStream(stream, clientUniqueCode, errch)

	go this.sendToStream(stream, clientUniqueCode, errch)

	return <-errch
}

func (this *server) receiveFromStream(stream pb.MyChatting_ChatServiceServer, clientUniqueCode int, errch chan error) {
	for {
		clientRequest, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving message from client :: %v", err)
			errch <- err
		} else {
			messageHandlerObject.Lock()

			messageHandlerObject.MessageQue = append(messageHandlerObject.MessageQue, message{
				ClientName:        clientRequest.Name,
				MessageBody:       clientRequest.Message,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode:  clientUniqueCode,
			})

			log.Printf("%v", messageHandlerObject.MessageQue[len(messageHandlerObject.MessageQue)-1])

			messageHandlerObject.Unlock()
		}
	}
}

func (this *server) sendToStream(stream pb.MyChatting_ChatServiceServer, clientUniqueCode int, errch chan error) {
	for {
		for {
			time.Sleep(500 * time.Millisecond)

			messageHandlerObject.Lock()

			if len(messageHandlerObject.MessageQue) == 0 {
				messageHandlerObject.Unlock()
				break
			}

			senderUniqueCode := messageHandlerObject.MessageQue[0].ClientUniqueCode
			senderNameForClient := messageHandlerObject.MessageQue[0].ClientName
			messageForClient := messageHandlerObject.MessageQue[0].MessageBody

			messageHandlerObject.Unlock()

			if senderUniqueCode != clientUniqueCode {

				err := stream.Send(&pb.ClientResponse{Name: senderNameForClient, Message: messageForClient})

				if err != nil {
					errch <- err
				}

				messageHandlerObject.Lock()

				if len(messageHandlerObject.MessageQue) > 1 {
					// Delete the message at index 0 after sending to receiver
					messageHandlerObject.MessageQue = messageHandlerObject.MessageQue[1:]
				} else {
					messageHandlerObject.MessageQue = []message{}
				}

				messageHandlerObject.Unlock()

			}

		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	// Init listener
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Could not listen to :5000 :: %v", err)
	}
	log.Println("Listening to :5000")

	// Create grpc server
	grpcserver := grpc.NewServer()

	// Register ChatService
	server := server{}
	pb.RegisterMyChattingServer(grpcserver, &server)

	// grpc listen and serve
	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Failed starting grpc server :: %v", err)
	}
}
