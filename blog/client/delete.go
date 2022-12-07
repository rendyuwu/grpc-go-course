package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting: %v", err)
	}

	log.Println("Blog was deleted!")
}
