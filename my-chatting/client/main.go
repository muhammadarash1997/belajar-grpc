package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/muhammadarash1997/my-chatting/pb"
	"google.golang.org/grpc"
)

type user struct {
	stream pb.MyChatting_ChatServiceClient
	name   string
}

func NewUser(stream pb.MyChatting_ChatServiceClient) *user {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed reading from console :: %v", err)
	}
	name = strings.Trim(name, "\r\n")

	return &user{stream, name}
}

func (this *user) sendMessage() {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed reading message from console :: %v", err)
		}
		message = strings.Trim(message, "\r\n")

		clientRequest := &pb.ClientRequest{
			Name:    this.name,
			Message: message,
		}

		err = this.stream.Send(clientRequest)
		if err != nil {
			log.Fatalf("Error sending message to request :: %v", err)
		}
	}
}

func (this *user) receiveMessage() {
	for {
		clientResponse, err := this.stream.Recv()
		if err != nil {
			log.Printf("Error receiving message from server :: %v", err)
		}
		fmt.Printf("%v: %v\n", clientResponse.GetName(), clientResponse.GetMessage())
	}
}

func main() {
	// Configurate server address
	serverAddress := configServer()

	// Create client connection
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to %v", serverAddress)
	}
	defer conn.Close()

	// Create grpc client
	client := pb.NewMyChattingClient(conn)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to grpc server :: %v", err)
	}

	user := NewUser(stream)

	go user.sendMessage()
	go user.receiveMessage()

	// Blocker
	blocker := make(chan bool)
	<-blocker
}

func configServer() string {
	fmt.Println("Enter Server Port :::")
	reader := bufio.NewReader(os.Stdin)
	serverAddress, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading from console :: %v", err)
	}
	serverAddress = strings.Trim(serverAddress, "\r\n")

	log.Printf("Connecting to %v", serverAddress)

	return serverAddress
}
