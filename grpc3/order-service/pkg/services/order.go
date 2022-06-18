package service

import (
    "context"
    "net/http"

    "order-service/pkg/client"
    "order-service/pkg/db"
    "order-service/pkg/models"
    "order-service/pkg/pb"
)

// Server API (Service)
type Server struct {
	pb.UnimplementedOrderServiceServer
    Handler          db.Handler
    ProductSvc client.ProductServiceClient
}

// Handler and Service and Repository
func (this *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    product, err := this.ProductSvc.FindOne(req.ProductId)
    if err != nil {
        return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
    } else if product.Status >= http.StatusNotFound {
        return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
    } else if product.Data.Stock < req.Quantity {
        return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Stock too less"}, nil
    }

    order := models.Order{
        Price:     product.Data.Price,
        ProductId: product.Data.Id,
        UserId:    req.UserId,
    }

    this.Handler.DB.Create(&order)

    res, err := this.ProductSvc.DecreaseStock(req.ProductId, order.Id)
    if err != nil {
        return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
    } else if res.Status == http.StatusConflict {
        this.Handler.DB.Delete(&models.Order{}, order.Id)

        return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
    }

    return &pb.CreateOrderResponse{
        Status: http.StatusCreated,
        Id:     order.Id,
    }, nil
}