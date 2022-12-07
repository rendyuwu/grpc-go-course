package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s", e.Message())
			log.Printf("Error code from server: %s", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("We probably send negative number")
			}
		} else {
			log.Fatalf("A non gRPC error: %v", err)
		}
	}

	log.Printf("Sqrt: %f", res.Result)
}
