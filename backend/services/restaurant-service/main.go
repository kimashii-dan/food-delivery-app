package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/pb"
	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/repository"
	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	db, err := repository.InitRestaurantServiceDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	restaurantRepo := repository.NewRestaurantRepository(db)
	menuItemRepo := repository.NewMenuItemRepository(db)

	restaurantService := service.NewRestaurantService(restaurantRepo, menuItemRepo)

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRestaurantServiceServer(grpcServer, restaurantService)
	reflection.Register(grpcServer)

	log.Printf("Restaurant service listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
