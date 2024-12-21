package main

import (
	"context"
	"log"
	"time"

	pb "github.com/utushkin/test_grpc/user" // Замените на ваш путь
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetUserRequest{Id: 1}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User: %s", res.Name)
}
