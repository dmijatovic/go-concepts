package main

import (
	"log"
	"os"
	"os/signal"

	grpcblog "dv4all/crud-modb-api/blog"
	"dv4all/crud-modb-api/server/modb"
)

func setLogFlags() {
	// show filename and line number in log when server crashes
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// listen for quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	// define logging
	setLogFlags()
	// connect to mongo
	modb.Connect()
	// start server in separate routing
	go grpcblog.NewGRPCSrv()
	// wait for quit signal
	<-quit
	//close mongoDB
	modb.Close()
	log.Println("Closing gRPC server...")
}
