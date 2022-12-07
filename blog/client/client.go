package main

import (
	"context"
	pb "github.com/rendyuwu/grpc-go-course/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "New Rendy",
		Title:    "A new title",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating %v", err)
	}

	log.Println("Blog was updated!")
}
