package main

import (
	"log"
	"net"

	pb "example.com/grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the Server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server Started at Port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
}
