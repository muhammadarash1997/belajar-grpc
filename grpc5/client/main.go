package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/muhammadarash1997/grpc5/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewWearableServiceClient(conn)

	stream, err := client.BeatsPerSecond(context.Background(), &pb.BeatsPerSecondRequest{Uuid: "mario"})
	if err != nil {
		log.Fatalln("Client error", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalln("Stream recv", err)
			}

			fmt.Printf("Second: %v Value: %v\n", resp.GetSecond(), resp.GetValue())
		}

	}()

	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("All done, possible error", stream.Context().Err())
			break
		}
	}
}
