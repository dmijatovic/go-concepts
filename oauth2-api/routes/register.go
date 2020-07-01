package routes

import (
	"net/http"
)

// Register all server routes with methods
func Register() *http.ServeMux {
	mux := http.NewServeMux()

	// demo page does do much
	mux.HandleFunc("/demo", demo())
	// users management page GET,POST,PUT, DELETE
	mux.HandleFunc("/users", handleUsers)
	// login issues JWT tokens
	mux.HandleFunc("/login", handleLogin)
	// verify user token
	mux.HandleFunc("/verify", handleVerify)

	//home page is static index.html
	mux.Handle("/", http.FileServer(http.Dir("./views/")))

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("API Home page"))
}

// func jsonHeader(res http.ResponseWriter) http.ResponseWriter {
// 	res.WriteHeader(http.StatusOK)
// 	res.Header().Set("Content-Type", "application/json")
// 	return res
// }
