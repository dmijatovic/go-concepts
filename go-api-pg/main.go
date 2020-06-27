package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"./api"
	"./db"
)

func onAppClose(done chan bool) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs

	log.Printf("Signal...%v", sig)

	done <- true
}

func main() {
	log.Printf("Starting oauth2 server")

	done := make(chan bool, 1)

	// connect to database
	myDb := db.Connect()

	// start https server
	go api.HandleRequests()

	//listen to close
	go onAppClose(done)

	// wait for appCloseEvent
	<-done
	//close db
	if myDb != nil {
		db.Close(myDb)
	}
	log.Println("Closing oauth2 server")
}
