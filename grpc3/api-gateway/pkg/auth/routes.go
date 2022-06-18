package auth

import (
	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/auth/routes"
	"api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func RegisterRoutes(router *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := router.Group("auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (this *ServiceClient) Register(ctx *gin.Context) {
	// Passing Context and AuthServiceClient into handler
	routes.Register(ctx, this.Client)
}

func (this *ServiceClient) Login(ctx *gin.Context) {
	// Passing Context and AuthServiceClient into handler
	routes.Login(ctx, this.Client)
}

//	ServiceClient {
//		AuthServiceClient {
//			ClientConnection,
//		}
//	}
