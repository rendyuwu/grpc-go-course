package main

import (
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.Calculator_PrimeServer) error {
	log.Printf("Prime function was invoked with %v", in)

	var divisor int64 = 2
	var number int64 = in.Number
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{Result: divisor})
			number /= divisor
		} else {
			divisor += 1
		}
	}

	return nil
}
