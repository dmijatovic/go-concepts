package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/credentials"

	deadline "../pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var host string = "localhost:50051"
var certFile string = "/home/dusan/test/go/grpc/ssl/cert/ca.crt"

type actionType struct {
	Type    string
	Payload string
}

// NOT SECURED - iNSECURE connection
func gRPCCnn() (*grpc.ClientConn, error) {
	cnn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Println("Client connection failed...END")
	}
	return cnn, err
}

// SECURE SSL connection
func gRPCTLSCnn() (*grpc.ClientConn, error) {
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalln("Failed to load SSL...", sslErr)
	}
	opts := grpc.WithTransportCredentials(creds)

	cnn, err := grpc.Dial(host, opts)
	if err != nil {
		log.Println("Client connection failed...END")
	}
	return cnn, err
}

func createClient(cnn *grpc.ClientConn) deadline.WithDeadlineClient {
	c := deadline.NewWithDeadlineClient(cnn)
	return c
}

func performRequest(grpcClient deadline.WithDeadlineClient, timeout time.Duration) {
	action := actionType{
		Type:    "NO_MATTER_WHAT",
		Payload: "Could be anything!",
	}

	// this defines the deadling of this request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// if cancel call cancel function
	defer cancel()

	res, err := grpcClient.Task(ctx, &deadline.Request{
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
	cnn, err := gRPCTLSCnn()
	if err != nil {
		log.Panicln(err)
	}
	// create client
	grpcClient := createClient(cnn)

	// first request within 1 second
	performRequest(grpcClient, 1*time.Second)
	// second withing 5 sec
	performRequest(grpcClient, 5*time.Second)
}
