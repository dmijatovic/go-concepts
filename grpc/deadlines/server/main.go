package main

import (
	"context"
	"log"
	"net"
	"time"

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

func (*grpcSrv) Task(ctx context.Context, req *deadline.Request) (*deadline.Response, error) {
	// action := actionStruct{
	// 	Type:    req.GetType(),
	// 	Payload: req.GetPayload(),
	// }

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
	// return custom error to client
	// return nil, status.Errorf(
	// 	codes.InvalidArgument,
	// 	fmt.Sprintf("Action type not supported: %v", action.Type),
	// )

	// after 3 seconds we return some response
	return &deadline.Response{
		Result: "After 3 seconds I am retruning you this result :-)",
	}, nil

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

	deadline.RegisterWithDeadlineServer(s, &grpcSrv{})

	log.Println("Starting gRPC server on...", host)

	log.Panic(s.Serve(lis))
}

func main() {
	// log.Println("Server starting")
	newGRPCSrv()
}
