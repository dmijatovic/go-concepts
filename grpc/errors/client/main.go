package main

import (
	"context"
	"log"

	errors "../pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var host string = "localhost:50051"

type actionType struct {
	Type    string
	Payload string
}

func gRPCCnn() (*grpc.ClientConn, error) {
	cnn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Println("Client connection failed...END")
	}
	return cnn, err
}

func createClient(cnn *grpc.ClientConn) errors.ErrorsClient {
	c := errors.NewErrorsClient(cnn)
	return c
}

func performRequest(grpcClient errors.ErrorsClient) {
	action := actionType{
		Type:    "NO_MATTER_WHAT",
		Payload: "Could be anything!",
	}
	res, err := grpcClient.Task(context.Background(), &errors.Request{
		Type:    action.Type,
		Payload: action.Payload,
	})

	if err != nil {
		// we need to test for custom gRPC errors from server
		re, ok := status.FromError(err)
		if ok {
			//this is gRPC error then we can extract
			// additional information like code and message
			// we could also test for code
			log.Printf("%v: %v", re.Code(), re.Message())
		} else {
			// not custom grpc error
			log.Printf(err.Error())
		}
	} else {
		log.Println("Response from server: ", res)
	}
}

func main() {
	log.Println("Starting client...")
	// create chat connection to server
	cnn, err := gRPCCnn()
	if err != nil {
		log.Panicln(err)
	}
	// create client
	grpcClient := createClient(cnn)
	// first send all messages and then receive responses
	performRequest(grpcClient)
}
