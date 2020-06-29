package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"../model"
)

func users() http.HandlerFunc {
	return handleUsers
}

func handleUsers(res http.ResponseWriter, req *http.Request) {
	// route based on request method (UPPER)
	upperMethod := strings.ToUpper(req.Method)
	switch upperMethod {
	case "GET":
		getAllUsers(res)
	case "POST":
		addNewUser(req, res)
	case "PUT":
		updateUser(req, res)
	case "PATCH":
		updateUser(req, res)
	case "DELETE":
		deleteUser(req, res)
	default:
		log.Printf("METHOD NOT SUPPORTED...%v", upperMethod)
		var data model.Response
		data.Status = http.StatusBadRequest
		data.StatusText = "Method not supported"
		data.Payload = "Method not supported"
		//return bad request
		json.NewEncoder(res).Encode(data)
	}

	// res.Write([]byte("This is users route"))
}

func getAllUsers(res http.ResponseWriter) {
	var data model.Response
	users, err := model.GetAllUsers()
	if err != nil {
		log.Println("getAllUsers: ", err)
		// data.Status = http.StatusInternalServerError
		// data.StatusText = "Internal server error"
		// data.Payload = err.Error()
		data = model.SetErrorResponse(err, model.ServerStatus{})
	} else {
		// data.Status = http.StatusOK
		// data.StatusText = "OK"
		// data.Payload = users
		data = model.SetOKResponse(users)
	}
	//write json header
	jsonHeader(res)
	json.NewEncoder(res).Encode(data)
}

func getUserFromReqBody(req *http.Request, res http.ResponseWriter) (model.User, error) {
	var data model.Response
	var user model.User
	//extract data from request body
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("getUserFromReqBody: ", err)

		data.Status = http.StatusBadRequest
		data.StatusText = "Failed to extract data from request.Body"
		data.Payload = err.Error()

		// return error
		json.NewEncoder(res).Encode(data)
		//exit
		return user, err
	}
	return user, nil
}

func addNewUser(req *http.Request, res http.ResponseWriter) {
	var data model.Response
	// extract user from body
	user, err := getUserFromReqBody(req, res)
	if err != nil {
		//exit on error
		return
	}
	// call insert statement
	id, err := model.AddNewUser(user)
	// check if insert failed
	if err != nil {
		data = model.SetErrorResponse(err, model.ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to extract data from request.Body",
		})
	} else {
		user.ID = id
		data = model.SetOKResponse(user)
	}

	json.NewEncoder(res).Encode(data)
}

func updateUser(req *http.Request, res http.ResponseWriter) {
	var data model.Response

	// extract user from body
	user, err := getUserFromReqBody(req, res)
	if err != nil {
		//exit on error
		return
	}

	data = model.SetOKResponse(user)

	json.NewEncoder(res).Encode(data)
}

func deleteUser(req *http.Request, res http.ResponseWriter) {
	var data model.Response

	// extract user from body
	user, err := getUserFromReqBody(req, res)
	if err != nil {
		//exit on error
		return
	}

	data = model.SetOKResponse(user)

	json.NewEncoder(res).Encode(data)
}
