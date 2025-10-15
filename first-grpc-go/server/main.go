package main

import (
	"context"
	pb "github.com/I-Frostbyte/first-grpc-go/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

// This code implements the service interface generated from our service definition in `helloworld.proto` and
// provisions a gRPC server and listens to requests on port 50051 from clients.

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "Hello World! ",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}