package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Resp is test response object
type resp struct {
	Status     int
	StatusText string
	Data       string
}

func main() {
	//define router
	router := mux.NewRouter()
	router.HandleFunc("/", testFn)
	//start http server and pass router handler
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// response := struct {
	// 	status     int
	// 	statusText string
	// 	data       string
	// }{status: 200, statusText: "OK", data: "This is test function from go api"}
	// fmt.Println(response)

	fmt.Println("Starting GO server on localhost:8080")
	// returns error when server crashes
	// oterwise it stays here
	err := srv.ListenAndServe()
	//it always returns non nil error
	log.Fatal(err)
}

func testFn(w http.ResponseWriter, r *http.Request) {
	//set json header
	w.Header().Set("Content-Type", "application/json")
	response := resp{Status: 200, StatusText: "OK", Data: "This is test function from go api"}
	fmt.Println(response)
	// map[string]bool{"ok": true}
	json.NewEncoder(w).Encode(response)
	// w.Write([]byte("This is test fn"))
}
