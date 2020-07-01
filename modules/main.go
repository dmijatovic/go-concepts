package main

import (
	"dv4all-module-demo/token"
	"log"
)

// import "./token"

func main() {
	println("Starting...")
	claims := token.CustomClaims{
		Profile: token.UserClaims{
			FirstName: "Dusan",
		},
	}

	cc, err := token.Sign(claims)
	if err != nil {
		log.Println("Failed to crate token...", err)
	}
	log.Println("Token...", cc)
}
