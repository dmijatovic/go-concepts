package routes

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"../token"
)

func handleVerify(res http.ResponseWriter, req *http.Request) {
	// route based on request method (UPPER)
	upperMethod := strings.ToUpper(req.Method)
	switch upperMethod {
	case "GET":
		verifyUser(req, res)
	case "POST":
		verifyUser(req, res)
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

func getTokenFromAuthHeader(req *http.Request) (string, error) {
	ah := req.Header.Get("Authorization")
	token := strings.Split(ah, " ")
	if len(token) == 2 {
		if strings.ToLower(token[0]) != "bearer" {
			return "", errors.New("Authorization header Bearer token missing")
		}
		return token[1], nil
	}
	return "", errors.New("Authorization header Bearer missing")
}

func verifyToken(tokenStr string) Response {
	valid, err := token.Verify(tokenStr)
	if valid == true {
		var r = struct {
			AccessToken string `json:"access_token"`
			TokenStatus string `json:"token_status"`
		}{
			AccessToken: tokenStr,
			TokenStatus: err,
		}
		return SetOKResponse(r)
	}
	return SetErrorResponse(errors.New(err), ServerStatus{
		Status:     http.StatusForbidden,
		StatusText: err,
	})
}

func verifyUser(req *http.Request, res http.ResponseWriter) {
	// var data Response
	tokenStr, err := getTokenFromAuthHeader(req)
	if err != nil {
		data := SetErrorResponse(err, ServerStatus{
			Status:     http.StatusForbidden,
			StatusText: "Forbidden",
		})
		data.ReturnResponse(res)
	} else {
		data := verifyToken(tokenStr)
		data.ReturnResponse(res)
	}
}
