package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "example.com/grpc/proto"
)

func CallSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming is Started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not Send Names %v", err)
	}

	receive_channel := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error Receiving String From the Server %v", err)
			}
			log.Println(message.Message)
		}
		close(receive_channel)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Names: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Could Not send the Name in Stream %v", err)
		}

		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-receive_channel

	log.Printf("Streaming Ended from Client")
}
