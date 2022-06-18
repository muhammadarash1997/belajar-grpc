package routes

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Model input
type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Here is where to communicate to other microservice
	res, err := c.Register(context.Background(), &pb.RegisterRequest{Email: body.Email, Password: body.Password})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
