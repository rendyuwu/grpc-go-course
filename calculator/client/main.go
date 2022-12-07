package main

import (
	pb "github.com/rendyuwu/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connet: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	//doSum(c)
	//doPrime(c)
	//doAverages(c)
	//doMax(c)
	//doSqrt(c, 10)
	doSqrt(c, -2)
}
