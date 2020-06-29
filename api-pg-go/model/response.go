package model

// Response type for json api response
// NOTE! json tags with "" and without spaces
type Response struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	// interface makes it possible
	// to send any format
	Payload interface{} `json:"payload"`
}
