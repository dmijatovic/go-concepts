package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"../pgdb"
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
	users, err := pgdb.GetAllUsers()
	if err != nil {
		data = SetErrorResponse(err, ServerStatus{})
	} else {
		data = SetOKResponse(users)
	}
	//write reponse
	data.ReturnResponse(res)
}

func getUserFromReqBody(req *http.Request, res http.ResponseWriter) (pgdb.InputUser, error) {
	var data Response
	var user pgdb.InputUser
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
	input, err := getUserFromReqBody(req, res)
	if err != nil {
		//exit on error
		return
	}
	// call insert statement
	user, err := pgdb.AddNewUser(input)
	// check if insert failed
	if err != nil {
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to extract data from request.Body",
		})
	} else {
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

	user, err := pgdb.UpdateUser(u)
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
	user, err := pgdb.DeleteUserByID(u.ID)
	if err != nil {
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
