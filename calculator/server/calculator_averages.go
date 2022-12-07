package main

import (
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Averages(stream pb.Calculator_AveragesServer) error {
	log.Println("Averages function was invoked")

	var sum int64 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{Result: float64(sum) / float64(count)})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		count += 1
		sum += req.Number
	}
}
