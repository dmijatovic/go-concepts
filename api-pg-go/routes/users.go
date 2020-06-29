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
		var data Response
		data.Status = http.StatusBadRequest
		data.StatusText = "Method not supported"
		data.Payload = "Method not supported"
		//return bad request
		data.ReturnResponse(res)
	}
}

func getAllUsers(res http.ResponseWriter) {
	var data Response
	users, err := model.GetAllUsers()
	if err != nil {
		data = SetErrorResponse(err, ServerStatus{})
	} else {
		data = SetOKResponse(users)
	}
	//write reponse
	data.ReturnResponse(res)
}

func getUserFromReqBody(req *http.Request, res http.ResponseWriter) (model.User, error) {
	var data Response
	var user model.User
	//extract data from request body
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("getUserFromReqBody: ", err)
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to extract data from request.Body",
		})
		data.ReturnResponse(res)
		//exit
		return user, err
	}
	return user, nil
}

func addNewUser(req *http.Request, res http.ResponseWriter) {
	var data Response
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
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to extract data from request.Body",
		})
	} else {
		user.ID = id
		data = SetOKResponse(user)
	}
	//write reponse
	data.ReturnResponse(res)
}

func updateUser(req *http.Request, res http.ResponseWriter) {
	var data Response

	// extract user from body
	u, err := getUserFromReqBody(req, res)
	if err != nil {
		//exit on error
		return
	}

	user, err := model.UpdateUser(u)
	if err != nil {
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to update user",
		})
	} else {
		data = SetOKResponse(user)
	}

	//write reponse
	data.ReturnResponse(res)
}

func deleteUser(req *http.Request, res http.ResponseWriter) {
	var data Response

	// extract user from body
	u, err := getUserFromReqBody(req, res)
	if err != nil {
		//exit on error
		return
	}
	user, err := model.DeleteUserByID(u.ID)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to delete user",
		})
	} else {
		data = SetOKResponse(user)
	}

	//write reponse
	data.ReturnResponse(res)
}
