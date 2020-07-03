package main

import (
	"log"
	"net"

	blog "../pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var host string = "0.0.0.0:50051"
var mongodb string = "mongodb://root:changeme@localhost:27017"
var mc mongo.Client

type grpcSrv struct{}

// create insecure gRPC server
func newGRPCSrv() {
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

func setLogFlags() {
	// show filename and line number in log when server crashes
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func mongoCnn() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb))
	if err != nil {
		return client, err
	}
	log.Println("Connected to mongodb...", mongodb)
	return client, nil
}

func main() {
	// define logging
	setLogFlags()
	// connect to mongo
	mc, err := mongoCnn()
	// log.Println(mc)
	if err != nil || mc == nil {
		// we panic
		log.Panic(err)
	}
	// start server
	newGRPCSrv()
}
