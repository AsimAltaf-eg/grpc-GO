package main

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc/proto"
)

func CallSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming Started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error While Client Streaming %v", err)
	}

	for _, name := range names.Names {

		req := &pb.HelloRequest{
			Names: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error While sending Request %v", err)
		}
		log.Printf("Sent Request with Name : %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error While sending Request %v", err)
	}
	log.Printf("%v", res.Messages)
}
