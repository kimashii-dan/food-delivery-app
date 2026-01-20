package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimashii-dan/food-delivery-app/backend/api/clients"
	"github.com/kimashii-dan/food-delivery-app/backend/api/handlers"
	"github.com/kimashii-dan/food-delivery-app/backend/api/middleware"
	"github.com/kimashii-dan/food-delivery-app/backend/pkg"
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

	// init grpc connection with restaurant service
	restaurantServicePort := os.Getenv("RESTAURANT_SERVICE_PORT")
	restaurantClient, restaurantConn := clients.NewRestaurantServiceClient(restaurantServicePort)
	defer restaurantConn.Close()

	// register handlers
	userHandler := handlers.NewUserHandler(userClient)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantClient)

	// init default web server
	r := gin.Default()

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtService := pkg.NewJWTService(jwtSecret)

	r.MaxMultipartMemory = 8 << 20

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// api endpoints
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
			users.POST("/logout", userHandler.Logout)
			users.POST("/refresh", userHandler.Refresh)
			users.GET("/me", middleware.CheckAuth(jwtService), userHandler.GetUser)
			users.POST("/addresses", middleware.CheckAuth(jwtService), userHandler.AddAddress)
			users.GET("/addresses", middleware.CheckAuth(jwtService), userHandler.GetAddresses)
		}

		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("", restaurantHandler.GetRestaurants)
			restaurants.GET("/:id", restaurantHandler.GetRestaurant)
			restaurants.GET("/:id/menu", restaurantHandler.GetMenu)
			restaurants.GET("/menu-items/:id", restaurantHandler.GetMenuItem)
			restaurants.GET("/:id/status", restaurantHandler.GetRestaurantStatus)
			restaurants.POST("/validate-items", restaurantHandler.ValidateMenuItems)
		}
	}

	// run server
	port := os.Getenv("API_PORT")
	log.Printf("API Gateway listening on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
