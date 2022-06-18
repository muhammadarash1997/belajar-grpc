package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Model  input
type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

// Handler
func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
