package main

import (
	"log"

	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order"
	"api-gateway/pkg/product"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	router := gin.Default()

	authSvc := *auth.RegisterRoutes(router, c)
	product.RegisterRoutes(router, c, &authSvc)
	order.RegisterRoutes(router, c, &authSvc)

	router.Run(c.Port)
}
