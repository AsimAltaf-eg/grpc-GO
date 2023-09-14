package main

import (
	"io"
	"log"

	pb "example.com/grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}

		if err == nil && req != nil {
			messages = append(messages, "-", req.Names)
		}

		log.Printf("Got Request with name: %v", req.Names)
	}
}
