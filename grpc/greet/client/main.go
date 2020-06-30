package main

import (
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
	log.Println("Client registered: ", c)
}
