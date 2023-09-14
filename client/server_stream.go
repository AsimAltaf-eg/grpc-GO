package main

import (
	"context"
	"io"
	"log"

	pb "example.com/grpc/proto"
)

func CallSayHelloServerStreaming(client pb.GreetServiceClient, Names *pb.NamesList) {

	stream, err := client.SayHelloServerStreaming(context.Background(), Names)

	if err != nil {
		log.Fatalf("Could not send Names %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error While Reading the Stream %v", err)
		}
		log.Printf("%s", message)
	}

	log.Printf("Streaming Fininshed")
}
