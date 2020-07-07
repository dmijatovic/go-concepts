package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

var host = ":8080"

func defineRoutes(l *log.Logger) *http.ServeMux {
	//create new router - mux
	router := http.NewServeMux()
	// create new route handler for hello page
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	// define routes
	router.Handle("/", hh)
	router.Handle("/goodbye", gb)

	return router
}

func createHTTPServer(host string, router *http.ServeMux, l *log.Logger) *http.Server {
	//create server
	api := &http.Server{
		Addr: host,
		// Handler can be anything that
		// implements ServeHTTP method
		Handler:      router,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return api
}

func main() {
	//create new logger with definitions
	l := log.New(os.Stdout, "myapp-api ", log.LstdFlags)
	//create routes
	r := defineRoutes(l)
	//create http sever instance
	api := createHTTPServer(host, r, l)

	// start server in its own routine
	// non-blocking
	go func() {
		// start server
		l.Println("Starting", host)
		// this isblocking
		log.Fatal(api.ListenAndServe())
	}()

	// register for os signals
	// listening for closing server
	stop := make(chan os.Signal)
	// broadcast message
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, os.Kill)

	// wait here to receive os signals
	sig := <-stop

	// close all systems
	l.Println("Terminate signal received:", sig)
	l.Println("Starting gracefull shutdown...")
	//define waiting for grafull shutdown to 30 secs
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	api.Shutdown(tc)
}
