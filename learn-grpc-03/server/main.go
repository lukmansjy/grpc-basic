package main

import (
	"context"
	"fmt"
	pb "github.com/lukmansjy/grpc-basic/learn-grpc-03/user"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) UnaryGetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:    "12312312",
		Name:  "Lukman",
		Email: request.Email, // ini hanya contoh
		Age:   17,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println(err.Error())
	}
}
