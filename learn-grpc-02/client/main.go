package main

import (
	"context"
	"fmt"
	pb "github.com/lukmansjy/grpc-basic/learn-grpc-02/proto/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"time"
)

func callerUnaryEcho(c pb.EchoClient, message string) {
	fmt.Println("--callerUnaryEcho--")
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		fmt.Println("error", err.Error())
		os.Exit(1)
	}
	fmt.Println(r.Message)
}

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("error", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	client := pb.NewEchoClient(conn)
	callerUnaryEcho(client, "message-1")
	callerUnaryEcho(client, "message-2")
}
