package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	helloworldpb "../pd/helloworld"
	hipb "../pd/hi"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func sendHello(helloworldC helloworldpb.HelloServiceClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := helloworldC.SayHello(ctx, &helloworldpb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("HelloService: %s", r.Message)
}

func sendHi(hiC hipb.HiServiceClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := hiC.SayHi(ctx, &hipb.HiRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("HiService: %s", r.Message)
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	helloworldC := helloworldpb.NewHelloServiceClient(conn)
	hiC := hipb.NewHiServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	
	sendHello(helloworldC, name)
	sendHi(hiC, name)
}
