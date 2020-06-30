package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greetpb "../pb"
	"google.golang.org/grpc"
)

type grpcSrv struct{}

// implement Greet interface from greetpb to our grpcSrv
func (*grpcSrv) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Greet server received request...", req)
	// get firstName from GreetRequest struct
	firstName := req.Greeting.GetFirstName()
	lastName := req.Greeting.GetLastName()
	// construct reponse
	msg := fmt.Sprintf("Hello %v %v!", firstName, lastName)
	// construct response
	resp := &greetpb.GreetResponse{
		Result: msg,
	}
	return resp, nil
}

func main() {
	println("Server works!")
	//default gRPC port on localhost
	host := "0.0.0.0:50051"
	//listen to tcp
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
	// create grpc server
	s := grpc.NewServer()
	// register our grpc service and pass grpcServer and out local grpc object
	// the Register method is autogenerated by protoc (see readme file)
	greetpb.RegisterGreetServiceServer(s, &grpcSrv{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve greetpb: %v", err)
	}
}