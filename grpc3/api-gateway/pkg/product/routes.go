package product

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/product/pb"
	"api-gateway/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func RegisterRoutes(router *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := router.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (this *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, this.Client)
}

func (this *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, this.Client)
}
