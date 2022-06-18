package routes

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Model input
type LoginRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Handler
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := LoginRequestBody{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Here is where to communicate to other microservice
	res, err := c.Login(context.Background(), &pb.LoginRequest{Email: body.Email, Password: body.Password})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}