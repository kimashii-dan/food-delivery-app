package clients

import (
	"log"

	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRestaurantServiceClient(address string) (pb.RestaurantServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to restaurant service: %v", err)
	}

	client := pb.NewRestaurantServiceClient(conn)
	return client, conn
}
