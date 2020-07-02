package main

import (
	"io"
	"log"
	"net"

	duplex "../pb"
	"google.golang.org/grpc"
)

// gRPC server
type grpcSrv struct{}
type actionStruct struct {
	Type    string
	Payload string
}

var host string = "0.0.0.0:50051"

// gRPC server needs to implement Chat method
func (*grpcSrv) Chat(stream duplex.ChatService_ChatServer) error {

	// liste to requests
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed by client")
			return nil
		}
		if err != nil {
			log.Println("Server error reading client stream: ", err)
			return err
		}
		// handle request
		result := handleRequests(req)

		log.Println("Sending result...", result)

		// send response
		err = stream.Send(result)
		if err != nil {
			log.Println("Failed to send response")
			return err
		}
	}
}

func handleRequests(req *duplex.Request) *duplex.Response {
	var resp duplex.Response
	action := actionStruct{
		Type:    req.GetType(),
		Payload: req.GetPayload(),
	}

	switch action.Type {
	case "GET_GREETING":
		resp.Result = "Hello " + action.Payload
		return &resp
	case "SAY_SOMETHING":
		resp.Result = "So, how are you today"
		return &resp
	default:
		resp.Result = "Action Type not supported..." + action.Type
		return &resp
	}
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

	duplex.RegisterChatServiceServer(s, &grpcSrv{})

	log.Println("Starting gRPC server on...", host)

	log.Panic(s.Serve(lis))
	// if err != nil {
	// 	log.Printf("Failed to register char service server")
	// }
}

func main() {
	// log.Println("Starting server...")
	//create new gRPC chat server
	newGRPCSrv()
}
