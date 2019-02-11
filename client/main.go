package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	helloworldPB "../pb/helloworld"
	hiPB "../pb/hi"
	crawlerPB "../pb/crawler"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func sendHello(helloworldConn helloworldPB.HelloServiceClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := helloworldConn.SayHello(ctx, &helloworldPB.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("HelloService: %s", r.Message)
}

func sendHi(hiConn hiPB.HiServiceClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := hiConn.SayHi(ctx, &hiPB.HiRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("HiService: %s", r.Message)
}

func getWeather(crawlerConn crawlerPB.CrawlerServiceClient, url string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := crawlerConn.GetWeather(ctx, &crawlerPB.WeatherRequest{Url: url})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("CrawlerService")
	for _, info := range r.Info {
		log.Printf("\t %s", info.Date)
		log.Printf("\t %s", info.Weather)
		log.Printf("\t %s", info.Temperature)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	helloworldConn := helloworldPB.NewHelloServiceClient(conn)
	hiConn := hiPB.NewHiServiceClient(conn)
	crawlerConn := crawlerPB.NewCrawlerServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	
	sendHello(helloworldConn, name)
	sendHi(hiConn, name)
	getWeather(crawlerConn, "https://www.cwb.gov.tw/V7/forecast/taiwan/inc/city/Taichung_City.htm")
}
