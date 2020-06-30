package token

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

var privateKey = []byte("ThisISMyPivateTestSigningKey12312312")

// UserClaims to be included in the JWT.
type UserClaims struct {
	ID        string `json:"id"`
	Roles     string `json:"roles"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// CustomClaims combines UserClaims and StandardJWT claims.
type CustomClaims struct {
	Profile UserClaims `json:"profile"`
	jwt.StandardClaims
}

// SetData for CustomClaims
func (cc *CustomClaims) SetData(d UserClaims) {
	cc.Profile = d
	cc.Audience = d.Roles
	cc.ExpiresAt = 60 * 60
	cc.Id = d.ID
	// cc.IssuedAt = time.Time.MarshalJSON
}

// Sign will create new token for valid user
// This method is used AFTER user is authenticated
// The data will be included in the token.
// DO NOT SEND secrity related information (like password) in the data
func Sign(claims CustomClaims) (string, error) {

	// var claims claimsType
	// // claims.data = data
	// claims.Audience = data.Roles
	// claims.ExpiresAt = 60 * 60
	// claims.Issuer = "dv4all"

	log.Println("clams...", claims)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := newToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func myKey(token *jwt.Token) (interface{}, error) {
	return privateKey, nil
}

// Verify will check if token is valid and return true/false and error message
func Verify(tokenStr string) (bool, string) {
	token, err := jwt.Parse(tokenStr, myKey)
	if token.Valid {
		return true, "Valid token"
	}
	// check for know errors
	ve, ok := err.(*jwt.ValidationError)
	if ok == true {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, "Token expired"
		}
		return false, "Invalid token"
	} else {
		return false, "Invalid token"
	}
}
