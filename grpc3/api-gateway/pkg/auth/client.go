package auth

import (
	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/config"
	"fmt"

	"google.golang.org/grpc"
)

// Create Client API
func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// Using WithInsecure() beacuse no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure()) // Create ClientConnection

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc) // Create AuthServiceClient
}
