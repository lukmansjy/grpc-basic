package main

import (
	"context"
	"fmt"
	pb "github.com/lukmansjy/grpc-basic/learn-grpc-02/proto/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("incoming request UnaryEcho")
	fmt.Printf("message: %s", in.Message)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "error UnaryEcho")
	}

	if v, ok := md["timestamp"]; ok {
		fmt.Println("metadata timestamp")
		for i, e := range v {
			fmt.Printf(" %d %s \n", i, e)
		}
	}
	return &pb.EchoResponse{Message: in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("error ", err.Error())
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Println("error ", err.Error())
		os.Exit(1)
	}
}
