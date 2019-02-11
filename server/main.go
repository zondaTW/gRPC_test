package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworldPB "../pb/helloworld"
	hiPB "../pb/hi"
	crawlerPB "../pb/crawler"
	weather "./crawler/weather"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.HelloServiceServer.
type server struct{}

// SayHello implements helloworld.HelloServiceServer
func (s *server) SayHello(ctx context.Context, in *helloworldPB.HelloRequest) (*helloworldPB.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &helloworldPB.HelloReply{Message: "Hello " + in.Name}, nil
}

// SayHi implements hi.HiServiceServer
func (s *server) SayHi(ctx context.Context, in *hiPB.HiRequest) (*hiPB.HiReply, error) {
	log.Printf("Received: %v", in.Name)
	return &hiPB.HiReply{Message: "Hi " + in.Name}, nil
}

func (s *server) GetWeather(ctx context.Context, in *crawlerPB.WeatherRequest) (*crawlerPB.WeatherReply, error) {
	log.Printf("Received: %v", in.Url)
	weatherInfoArray := weather.GetWeatherInfo(in.Url)
	return &crawlerPB.WeatherReply{Info: weatherInfoArray}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworldPB.RegisterHelloServiceServer(s, &server{})
	hiPB.RegisterHiServiceServer(s, &server{})
	crawlerPB.RegisterCrawlerServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}