package product

import (
	"fmt"
	"google.golang.org/grpc"
	"api-gateway/pkg/config"
	"api-gateway/pkg/product/pb"
)

// Create Client API
func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}