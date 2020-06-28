package main

import (
	"log"
	"net/http"
	"time"

	"./router"
)

func main() {
	println("It works")

	rt := router.Router{}

	api := &http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// Handler can be anything that
		// implements ServeHTTP method
		Handler: rt,
	}
	log.Fatal(api.ListenAndServe())
}
