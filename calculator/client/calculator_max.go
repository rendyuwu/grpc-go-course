package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"io"
	"log"
	"sync"
	"time"
)

func doMax(c pb.CalculatorClient) {
	log.Println("doMax function was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	numbers := []int64{9, 1, 4, 5, 3, 8, 3, 7, 11, 15, 13, 19}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, number := range numbers {
			log.Printf("Send request: %d", number)
			if err := stream.Send(&pb.MaxRequest{Number: number}); err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error closing stream: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received a new maximum: %d", res.Result)
		}
	}()

	wg.Wait()
}
