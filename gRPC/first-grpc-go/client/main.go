package main

import (
	"context"
	pb "github.com/I-Frostbyte/first-grpc-go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// This code creates a gRPC channel to communicate with the gRPC server, creates a
// gRPC stub (client) to perform RPCs, and finally calls our SayHello function on the gRPC server.

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloWorldRequest{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}

	log.Printf("Response from gRPC server's SayHello function: %s", r.GetMessage())
}