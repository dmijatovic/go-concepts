package routes

import (
	"encoding/json"
	"net/http"

	"../model"
)

func demo() http.HandlerFunc {
	return handleDemo
}

func handleDemo(resp http.ResponseWriter, req *http.Request) {
	data := model.Response{
		Status:     http.StatusOK,
		StatusText: "OK",
		Payload:    "This is demo response",
	}
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(data)
}
