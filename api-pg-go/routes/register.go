package routes

import (
	"net/http"
)

// Register all server routes with methods
func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/demo", demo())
	mux.HandleFunc("/users", handleUsers)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("API Home page"))
}

func jsonHeader(resp http.ResponseWriter) http.ResponseWriter {
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	return resp
}
