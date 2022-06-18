package main

import (
	"fmt"
	"log"
	"net"

	"auth-service/pkg/config"
	"auth-service/pkg/db"
	"auth-service/pkg/pb"
	"auth-service/pkg/services"
	"auth-service/pkg/utils"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	handler := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	// Create Listener
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	fmt.Println("Auth Svc on", c.Port)

	// Create Server API
	s := services.Server{
		Handler:   handler,
		Jwt: jwt,
	}

	// Create gRPC Server which still has no services
	grpcServer := grpc.NewServer()

	// Register Server API (Service) into gRPC Server and then the gRPC Server will has services from Server API (Service)
	pb.RegisterAuthServiceServer(grpcServer, &s)

	// Serve accepts incoming connections on the Listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
