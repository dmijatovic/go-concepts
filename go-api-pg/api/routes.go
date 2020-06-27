package api

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

func patchUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "patchUser endpoint")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "loginUser endpoint")
}

func verifyUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "verifyUser endpoint")
}
