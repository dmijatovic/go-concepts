package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../pgdb"
)

// func addHeader(*w *http.ResponseWriter) {
// 	w.Header().Set("Content-Type", "application/json")
// }

func allUsers(w http.ResponseWriter, r *http.Request) {
	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-server", "oauth2 Golang demo")
	//declare user struct
	var users []pgdb.User
	var err error
	//get all users from db
	users, err = pgdb.AllUsers(nil)
	// convert to json
	// json, e := json.Marshal(users)
	// if err != nil || e != nil {
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// log.Printf("Users...%v", users)
		// write json reponse
		// w.Write(json)
		// fmt.Fprintf(w, "All users endpoint: %v", users)
		// write json response with users
		json.NewEncoder(w).Encode(users)
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var u pgdb.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "newUser %v", u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var u pgdb.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "deleteUser %v", u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var u pgdb.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "updateUser %v", u)
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	var u pgdb.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "patchUser %v", u)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var u pgdb.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "loginUser %v", u)
}

func verifyUser(w http.ResponseWriter, r *http.Request) {
	var u pgdb.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		//send error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "verifyUser %v", u)
}
