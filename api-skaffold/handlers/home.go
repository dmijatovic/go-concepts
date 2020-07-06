package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

//HomeHandler handles home route requests.
type HomeHandler struct {
	logger *log.Logger
}

//NewHomeHandler creates HomeHandler for /
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		log.Println("Home GET request")
		getHome(rw, req)
	case http.MethodPost:
		log.Println("Home POST request")
		postHome(rw, req)
	default:
		log.Println("Home not supported method:", req.Method)
		notSupported(rw, req)
	}
}

func getHome(rw http.ResponseWriter, req *http.Request) {
	data := apiResponse{
		Status:     http.StatusOK,
		StatusText: "OK",
		Payload:    "This is my payload response on GET HOME",
	}
	json.NewEncoder(rw).Encode(data)
}

func postHome(rw http.ResponseWriter, req *http.Request) {
	data := apiResponse{
		Status:     http.StatusOK,
		StatusText: "OK",
		Payload:    "This is my payload response on POST HOME",
	}
	json.NewEncoder(rw).Encode(data)
}

func notSupported(rw http.ResponseWriter, req *http.Request) {
	data := apiResponse{
		Status:     http.StatusNotImplemented,
		StatusText: "Not implemented",
		Payload:    "The REST method is not supported",
	}
	json.NewEncoder(rw).Encode(data)
}
