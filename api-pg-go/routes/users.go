package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

func users() http.HandlerFunc {
	return handleUsers
}

func handleUsers(resp http.ResponseWriter, req *http.Request) {
	var data model.Response
	users, err := model.GetAllUsers()
	if err != nil {
		log.Println(err)
		data.Status = http.StatusInternalServerError
		data.StatusText = "Internal server error"
		data.Payload = err.Error()
	} else {
		data.Status = http.StatusOK
		data.StatusText = "OK"
		data.Payload = users
	}
	//write json header
	hresp := jsonHeader(resp)
	json.NewEncoder(hresp).Encode(data)
	// resp.Write([]byte("This is users route"))
}
