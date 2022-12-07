package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorClient) {
	log.Println("doSum was invoked")

	request := &pb.SumRequest{
		A: 10,
		B: 5,
	}

	res, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result of %d + %d = %d", request.A, request.B, res.Result)
}
