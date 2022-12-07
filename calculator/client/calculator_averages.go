package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"log"
	"time"
)

func doAverages(c pb.CalculatorClient) {
	log.Println("doAverages function was invoked")

	reqs := []*pb.AverageRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.Averages(context.Background())
	if err != nil {
		log.Printf("Error while calling Averages: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Averages: %v", err)
	}

	log.Printf("Averages: %f", res.Result)
}
