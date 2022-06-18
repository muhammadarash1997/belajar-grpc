package order

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func RegisterRoutes(router *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := router.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (this *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, this.Client)
}