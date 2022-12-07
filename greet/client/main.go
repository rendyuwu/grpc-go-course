package main

import (
	pb "github.com/rendyuwu/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

var addr string = "0.0.0.0:50051"

func main() {
	tls := true // change that to false if needed
	var opts []grpc.DialOption

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "localhost")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connet: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	//doGreetWithDeadline(c, 5*time.Second)
	//doGreetWithDeadline(c, 1*time.Second)
}
