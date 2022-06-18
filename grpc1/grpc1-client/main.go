package main

import (
	"context"
	"flag"
	"log"
	"time"

	"grpc1-client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// Create Client Connection
	cc, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	// Create GreeterClient
	client := pb.NewChatServiceClient(cc)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Here is where to communicate to other microservice
	res, err := client.SayHello(ctx, &pb.RequestMessage{Body: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.GetBody())
}
