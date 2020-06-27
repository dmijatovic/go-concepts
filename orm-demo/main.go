package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func startServer(router *mux.Router) {
	//start http server and pass router handler
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting GO server on localhost:8080")
	// returns error when server crashes
	// oterwise it stays here
	err := srv.ListenAndServe()
	//it always returns non nil error
	log.Fatal(err)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", helloWorld).Methods("GET")

	router.HandleFunc("/login", loginUser).Methods("POST")
	router.HandleFunc("/verify", verifyUser).Methods("POST")

	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users", newUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	router.HandleFunc("/users", patchUser).Methods("PATCH")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")

	startServer(router)
}

func main() {
	fmt.Println("Go ORM api demo")
	handleRequests()
}
