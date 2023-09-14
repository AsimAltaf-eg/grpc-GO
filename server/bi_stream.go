package main

import (
	"io"
	"log"
	"time"

	pb "example.com/grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("Streaming Ended from Client")
			return nil
		}

		if err == nil && req != nil {
			log.Printf("Got Request with Name : %v", req.Names)
			res := &pb.HelloResponse{
				Message: "Hello! - " + req.Names,
			}

			if err := stream.Send(res); err != nil {
				log.Fatalf("Error Sending the Response to the Client %v", err)
				return err
			}
		}

		if err != nil {
			log.Fatalf("Error Receiving the Stream from Client %v", err)
		}

		time.Sleep(2 * time.Second)
	}
}
