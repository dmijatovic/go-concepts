package main

import (
	"fmt"
	"net/http"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users endpoint")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "newUser endpoint")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleteUser endpoint")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateUser endpoint")
}
