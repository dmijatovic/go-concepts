package main

import (
	"context"
	"log"

	blog "../pb"
	"google.golang.org/grpc"
)

var host string = "localhost:50051"

func performRequest(grpcClient blog.BlogSvcClient) {

	res, err := grpcClient.Create(context.Background(), &blog.Blog{
		Id:       "1",
		AuthorId: "Author id",
		Title:    "Blog title",
		Content:  "Blog content",
	})
	if err != nil {
		log.Printf(err.Error())
	}
	log.Println(res)
}

func gRPCCnn() (*grpc.ClientConn, error) {
	cnn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Println("Client connection failed...END")
	}
	return cnn, err
}

func createClient(cnn *grpc.ClientConn) blog.BlogSvcClient {
	c := blog.NewBlogSvcClient(cnn)
	return c
}

func main() {
	log.Println("Starting client")

	cnn, err := gRPCCnn()
	if err != nil {
		log.Panicln(err)
	}
	// create client
	grpcClient := createClient(cnn)
	// first send all messages and then receive responses
	performRequest(grpcClient)

}
