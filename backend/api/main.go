package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimashii-dan/food-delivery-app/api/clients"
	"github.com/kimashii-dan/food-delivery-app/api/handlers"
)

func main() {
	// load all .env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// init grpc connection with user service
	userServicePort := os.Getenv("USER_SERVICE_PORT")
	userClient, userConn := clients.NewUserServiceClient(userServicePort)
	defer userConn.Close()

	// register user handlers
	userHandler := handlers.NewUserHandler(userClient)

	// init default web server
	r := gin.Default()

	// api endpoints
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
			users.POST("/refresh", userHandler.Refresh)
		}
	}

	// run server
	port := os.Getenv("API_PORT")
	log.Printf("API Gateway listening on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
