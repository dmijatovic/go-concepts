package model

import "net/http"

// ServerStatus is send as part of api response
// this is quite similair how axios responds
type ServerStatus struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

// Response type for json api response
// NOTE! json tags with "" and without spaces
type Response struct {
	ServerStatus
	// interface makes it possible
	// to send any format
	Payload interface{} `json:"payload"`
}

// SetErrorResponse creates standardized api response
func SetErrorResponse(err error, state ServerStatus) Response {
	if state.Status == 0 {
		state.Status = http.StatusInternalServerError
	}
	if state.StatusText == "" {
		state.StatusText = "Internal Server Error"
	}
	var r Response

	r.Status = state.Status
	r.StatusText = state.StatusText
	r.Payload = err.Error()

	return r
}

// SetOKResponse creates standard response with status and payload
func SetOKResponse(data interface{}) Response {
	var r Response

	r.Status = http.StatusOK
	r.StatusText = "OK"
	r.Payload = data

	return r
}
