package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.CalculatorClient) {
	log.Println("doPrime function was invoked")

	stream, err := c.Prime(context.Background(), &pb.PrimeRequest{Number: 1203242})
	if err != nil {
		log.Fatalf("Error while calling Prime: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}

		log.Printf("Result %d", msg.Result)
	}
}
