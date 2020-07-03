package grpcblog

import (
	"context"
	"log"
	"net"

	"dv4all/crud-modb-api/modb"

	blog "../pb"

	"google.golang.org/grpc"
)

// Host gRPC server
var host string = "0.0.0.0:50051"

type grpcSrv struct{}

// type jsonStr json.Encoder

// func (b *jsonStr) Write(str []byte) (n int, e error) {
// 	log.Println(str)
// 	b = string(str)
// 	return 20, nil
// }

// Create new blogpost
func (*grpcSrv) Create(ctx context.Context, b *blog.Blog) (*blog.Response, error) {

	log.Println("Create blog request")

	// json.NewEncoder(jsonStr).Encode(b)
	// item, err := modb.CreateBlog(modb.BlogItem{
	// 	AuthorID: b.GetAuthorId(),
	// 	Title:    b.GetTitle(),
	// 	Content:  b.GetContent(),
	// })
	item, err := modb.CreateBlog(b)

	if err != nil {
		return nil, err
	}

	log.Println("Created blog...", item)

	return &blog.Response{
		Status:     200,
		StatusText: "OK",
		Payload:    []byte("This is my string to bytes"),
	}, nil

}

// NewGRPCSrv creates insecure gRPC server
func NewGRPCSrv() {
	// start listening on tcp port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Panicf("gRPC server failed: %v", err)
	}
	defer lis.Close()

	// create grpc server
	s := grpc.NewServer()
	defer s.Stop()

	blog.RegisterBlogSvcServer(s, &grpcSrv{})

	log.Println("Starting gRPC server on...", host)
	// ch <- fmt.Sprintf("Starting gRPC server on...%v", host)

	log.Panic(s.Serve(lis))
}
