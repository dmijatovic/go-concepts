package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"../pgdb"
)

// LoginCredentials of user
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	// route based on request method (UPPER)
	upperMethod := strings.ToUpper(req.Method)
	switch upperMethod {
	case "POST":
		loginUser(req, res)
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

func getCredentialsFromBody(req *http.Request, res http.ResponseWriter) (LoginCredentials, error) {
	var data Response
	var cred LoginCredentials

	err := json.NewDecoder(req.Body).Decode(&cred)
	if err != nil {
		log.Println("getUserFromReqBody: ", err)
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusBadRequest,
			StatusText: "Failed to extract data from request.Body",
		})
		data.ReturnResponse(res)
		//exit
		return cred, err
	}
	return cred, nil
}

func validPassword(loginPass string, userPass string) bool {

	log.Println(loginPass)
	log.Println(userPass)

	if loginPass == userPass {
		return true
	} else {
		return false
	}
}

func signToken(user pgdb.User) string {
	token := fmt.Sprintf("user %v", user)
	return token
}

func loginUser(req *http.Request, res http.ResponseWriter) {
	var data Response

	cred, err := getCredentialsFromBody(req, res)
	if err != nil {
		return
	}

	data = SetOKResponse(cred)

	user, err := pgdb.GetUserByEmail(cred.Email)
	if err != nil {
		data = SetErrorResponse(err, ServerStatus{
			Status:     http.StatusForbidden,
			StatusText: "User email or password incorrect",
		})
		data.ReturnResponse(res)
		return
	}

	if validPassword(cred.Password, user.Password) == true {
		data = SetOKResponse(user)
		// data = SetOKResponse(signToken(user))
	} else {
		data.Status = http.StatusForbidden
		data.StatusText = "Forbidden"
		data.Payload = "User email or password incorrect"
	}
	// // send response
	data.ReturnResponse(res)
}
