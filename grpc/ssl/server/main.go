package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/credentials"

	deadline "../pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPC server
type grpcSrv struct{}
type actionStruct struct {
	Type    string
	Payload string
}

var host string = "0.0.0.0:50051"
var certFile string = "/home/dusan/test/go/concepts/grpc/ssl/cert/server.crt"
var keyFile string = "/home/dusan/test/go/concepts/grpc/ssl/cert/server.pem"

func (*grpcSrv) Task(ctx context.Context, req *deadline.Request) (*deadline.Response, error) {
	// we will wait for 3 seconds and check for
	// cancelation each second
	for i := 0; i < 3; i++ {
		// check if client cancelled request
		// using context (see client) for deadline implementation
		if ctx.Err() == context.Canceled {
			log.Println("Client cancelled request")
			return nil, status.Errorf(codes.Canceled, "The client cancelled request")
		}
		time.Sleep(1 * time.Second)
	}

	// after 3 seconds we return some response
	return &deadline.Response{
		Result: "After 3 seconds I am retruning you this result :-)",
	}, nil

}

// create insecure gRPC server
func newGRPCSrv() {
	// start listening on tcp port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Panicf("gRPC server failed: %v", err)
	}
	// create grpc server
	s := grpc.NewServer()

	deadline.RegisterWithDeadlineServer(s, &grpcSrv{})

	log.Println("Starting gRPC server on...", host)

	log.Panic(s.Serve(lis))
}

// create SECURE gRPC server
func newGRPCTLSrv() {
	// start listening on tcp port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Panicf("gRPC server failed: %v", err)
	}
	// SSL setup
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalln("Failed to load SSL...", sslErr)
	}
	opts := grpc.Creds(creds)

	// create grpc server
	s := grpc.NewServer(opts)

	deadline.RegisterWithDeadlineServer(s, &grpcSrv{})

	log.Println("Starting gRPC server on...", host)

	log.Panic(s.Serve(lis))
}

func main() {
	// log.Println("Server starting")
	// newGRPCSrv()
	newGRPCTLSrv()
}
