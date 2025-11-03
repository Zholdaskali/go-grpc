package main

import (
	"github/Zholdaskali/go-grpc/pkg/api/example"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	service := ExampleService{}
	example.RegisterExampleServer(server, service)
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server started: 8080")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type ExampleService struct {
	example.UnimplementedExampleServer
}
