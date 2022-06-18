package client

import (
	"context"
	"fmt"
	"order-service/pkg/pb"

	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

// Create Client API
func InitProductServiceClient(url string) ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (this *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return this.Client.FindOne(context.Background(), req)
}

func (this *ProductServiceClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return this.Client.DecreaseStock(context.Background(), req)
}
