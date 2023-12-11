package main

import (
	"context"
	"fmt"
	pb "github.com/lukmansjy/grpc-basic/learn-grpc-03/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}

func isValidAppKey(appKey []string) bool {
	if len(appKey) < 1 {
		return false
	}

	return appKey[0] == "1234545"
}

func (s *server) UnaryGetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:    "12312312",
		Name:  "Lukman",
		Email: request.Email, // ini hanya contoh
		Age:   17,
	}, nil
}

func unaryInterceptorImpl(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Printf("Incomming request: %s \n", info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid Argument")
	}

	if !isValidAppKey(md["app_key"]) {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	m, err := handler(ctx, req)
	if err != nil {
		log.Fatal("error in grpc service", err)
	}
	return m, err
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err.Error())
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptorImpl))
	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println(err.Error())
	}
}
