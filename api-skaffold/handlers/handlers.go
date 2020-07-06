package handlers

import (
	"net/http"
)

type apiResponse struct {
	Status     int         `json:"status"`
	StatusText string      `json:"statusText"`
	Payload    interface{} `json:"payload"`
}

// Setup route handlers for this api.
func Setup() *http.ServeMux {

	mux := http.NewServeMux()
	mux.Handle("/", NewHomeHandler())
	return mux
}
