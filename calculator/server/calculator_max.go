package main

import (
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.Calculator_MaxServer) error {
	log.Println("Max function was invoked")

	var maximum int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			if err := stream.Send(&pb.MaxResponse{Result: maximum}); err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
			}
		}
	}
}
