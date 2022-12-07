package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v", err)
	}

	log.Printf("Blog was read: %v", res)
	return res
}
