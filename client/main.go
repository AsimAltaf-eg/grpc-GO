package main

import (
	"log"

	pb "example.com/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // insecure dialing to connect with the server

	if err != nil {
		log.Fatalf("Failed to Connect to the server %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"Asim", "Rizwan", "Zafeer", "Talha", "Anas"},
	}

	//CallSayHello(client)
	//CallSayHelloServerStreaming(client, names)
	//CallSayHelloClientStreaming(client, names)
	CallSayHelloBidirectionalStreaming(client, names)
}
