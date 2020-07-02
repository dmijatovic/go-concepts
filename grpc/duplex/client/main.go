package main

import (
	"context"
	"io"
	"log"
	"time"

	duplex "../pb"
	"google.golang.org/grpc"
)

var host string = "localhost:50051"

// type grpcCnn struct{}
type chatSvc struct{}

type action struct {
	Type    string
	Payload string
}

var msgToSend = []action{{
	Type:    "TRAY",
	Payload: "No idea how this works",
}, action{
	Type:    "GET_HELLO",
	Payload: "Duenas",
}, action{
	Type:    "GET_GREETING",
	Payload: "Mijatovic",
}}

func sendMessages(stream duplex.ChatService_ChatClient) {
	// send all messages in message buffer
	for i, msg := range msgToSend {
		log.Println("Sending message...", i)

		err := stream.Send(&duplex.Request{
			Type:    msg.Type,
			Payload: msg.Payload,
		})
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
	// send end signal
	stream.CloseSend()
}

func listenForServerResponse(stream duplex.ChatService_ChatClient) {
	log.Println("Listen for server messages...START")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of server response stream...")
			break
		}
		if err != nil {
			log.Println(err)
		}
		log.Println("Response: ", msg.GetResult())
	}
	log.Println("Listen for server messages...END")
}

func sequentialProcess(grpcClient duplex.ChatServiceClient) {
	// create stream
	stream, err := grpcClient.Chat(context.Background())
	if err != nil {
		log.Println("Failed to create stream...", err)
	}
	//send messages
	sendMessages(stream)
	//receive responses
	listenForServerResponse(stream)
}

func parallelProcess(grpcClient duplex.ChatServiceClient) {
	// create stream
	stream, err := grpcClient.Chat(context.Background())
	if err != nil {
		log.Println("Failed to create stream...", err)
	}

	stopSignal := make(chan struct{})

	// parallel go routine
	go func(stream duplex.ChatService_ChatClient) {
		sendMessages(stream)
	}(stream)

	go func(stream duplex.ChatService_ChatClient) {
		// propgram is `stuck` here
		// untill the server breaks
		listenForServerResponse(stream)
		close(stopSignal)
	}(stream)

	<-stopSignal
}

func gRPCCnn() (*grpc.ClientConn, error) {
	cnn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Println("Client connection failed...END")
	}
	return cnn, err
}

func createClient(cnn *grpc.ClientConn) duplex.ChatServiceClient {
	c := duplex.NewChatServiceClient(cnn)
	return c
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
	sequentialProcess(grpcClient)
	// parallel send/receive using fo routing
	parallelProcess(grpcClient)
}
