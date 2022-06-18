package services

import (
	"context"
	"net/http"
	"product-service/pkg/db"
	"product-service/pkg/models"
	"product-service/pkg/pb"
)

// Server API (Service)
type Server struct {
	pb.UnimplementedProductServiceServer
	Handler db.Handler
}

// Handler and Service and Repository
func (this *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	result := this.Handler.DB.Create(&product)

	if result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
        Status: http.StatusCreated,
        Id:     product.Id,
    }, nil
}

// Handler and Service and Repository
func (this *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	result := this.Handler.DB.First(&product, req.Id)

	if result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

// Handler and Service and Repository
func (this *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := this.Handler.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := this.Handler.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	this.Handler.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	this.Handler.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
