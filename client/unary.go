package main

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could Not Greet %v", err)
	}
	log.Printf("Server Says: %s", res.Message)
}
