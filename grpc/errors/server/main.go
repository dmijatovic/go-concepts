package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	errors "../pb"
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

func (*grpcSrv) Task(ctx context.Context, req *errors.Request) (*errors.Response, error) {
	action := actionStruct{
		Type:    req.GetType(),
		Payload: req.GetPayload(),
	}

	//sleep for second
	time.Sleep(1 * time.Second)

	//return custom error to client
	return nil, status.Errorf(
		codes.InvalidArgument,
		fmt.Sprintf("Action type not supported: %v", action.Type),
	)

}

// create gRPC server
func newGRPCSrv() {
	// start listening on tcp port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Panicf("gRPC server failed: %v", err)
	}
	// create grpc server
	s := grpc.NewServer()

	errors.RegisterErrorsServer(s, &grpcSrv{})

	log.Println("Starting gRPC server on...", host)

	log.Panic(s.Serve(lis))
}

func main() {
	// log.Println("Server starting")
	newGRPCSrv()
}
