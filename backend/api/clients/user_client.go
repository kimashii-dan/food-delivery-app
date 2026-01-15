package clients

import (
	"log"

	"github.com/kimashii-dan/food-delivery-app/services/user-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient(address string) (pb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	client := pb.NewUserServiceClient(conn)
	return client, conn
}
