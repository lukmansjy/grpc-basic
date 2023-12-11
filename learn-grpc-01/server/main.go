package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/lukmansjy/grpc-basic/learn-grpc-01/student"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	fmt.Println("Incoming request for FindStudentByEmail")
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := os.ReadFile("data/students.json")
	if err != nil {
		log.Fatalf("Error read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalf("Error unmarshal data json", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error server grpc", err.Error())
	}
}
