package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Rendy"},
		{FirstName: "Wijaya"},
		{FirstName: "Budi"},
	}

	waitC := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v", req)
			if err := stream.Send(req); err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error closing stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v", res.Result)
		}

		close(waitC)
	}()

	<-waitC
}
