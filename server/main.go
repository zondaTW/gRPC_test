package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	helloworldpb "../pd/helloworld"
	hipb "../pd/hi"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.HelloServiceServer.
type server struct{}

// SayHello implements helloworld.HelloServiceServer
func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &helloworldpb.HelloReply{Message: "Hello " + in.Name}, nil
}

// SayHi implements hi.HiServiceServer
func (s *server) SayHi(ctx context.Context, in *hipb.HiRequest) (*hipb.HiReply, error) {
	log.Printf("Received: %v", in.Name)
	return &hipb.HiReply{Message: "Hi " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworldpb.RegisterHelloServiceServer(s, &server{})
	hipb.RegisterHiServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}