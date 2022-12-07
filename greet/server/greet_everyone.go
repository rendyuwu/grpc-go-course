package main

import (
	pb "github.com/rendyuwu/grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res := "Hello " + req.FirstName + "!"
		if err := stream.Send(&pb.GreetResponse{Result: res}); err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}
	}
}
