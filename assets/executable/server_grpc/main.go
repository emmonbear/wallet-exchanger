package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"log"
	"net"

	pb "github.com/utushkin/test_grpc/user" // Замените на ваш путь
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Пример данных
	users := map[int32]string{
		1: "John Doe",
		2: "Jane Smith",
	}

	name, exists := users[req.Id]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "User not found")
	}

	return &pb.GetUserResponse{Id: req.Id, Name: name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
