package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/kimashii-dan/food-delivery-app/backend/services/user-service/pb"
	"github.com/kimashii-dan/food-delivery-app/backend/services/user-service/repository"
	"github.com/kimashii-dan/food-delivery-app/backend/services/user-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	db, err := repository.Init(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	addressRepo := repository.NewAddressRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")

	jwtService := service.NewJWTService(jwtSecret)
	userService := service.NewUserService(userRepo, addressRepo, jwtService)

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)
	reflection.Register(grpcServer)

	log.Printf("User service listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
