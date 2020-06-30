package main

import (
	"context"
	"log"

	greetpb "../pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client initalizing...")
	// establish connection (without SSL)
	cnn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %v", err)
	}
	// close at the end
	defer cnn.Close()
	// register client for greetpb service
	c := greetpb.NewGreetServiceClient(cnn)
	unaryCnn(c)
}

func unaryCnn(c greetpb.GreetServiceClient) {
	log.Println("Starting Unary connection...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dusan",
			LastName:  "Mijatovic",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Client Greet error: %f", err)
	}
	log.Println("Server response...", res.Result)
}
