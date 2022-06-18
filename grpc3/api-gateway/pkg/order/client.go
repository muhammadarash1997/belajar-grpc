package order

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/order/pb"
	"fmt"

	"google.golang.org/grpc"
)

// Create Client API
func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// Using WithInsecure() beacuse no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())  // Create ClientConnection

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc) // Create OrderServiceClient
}
