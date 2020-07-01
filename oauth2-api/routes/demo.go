package routes

import (
	"net/http"
)

func demo() http.HandlerFunc {
	return handleDemo
}

func handleDemo(res http.ResponseWriter, req *http.Request) {
	data := SetOKResponse("This is demo response")
	data.ReturnResponse(res)
}
