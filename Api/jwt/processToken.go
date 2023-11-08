package jwt

import (
	"errors"
	"strings"
    "net/http"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

type claim struct {
	jwt.RegisteredClaims
}

var Auth bool

func ProcessToken(tk string, JWTSign string) (*claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims claim
	
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "invalid token format", errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if !tkn.Valid {
		return &claims, false, "invalid token", errors.New("invalid token")
	}

	return &claims, true, tk, err
}

func ValidateUser(ctx *gin.Context, request *http.Request) (bool, int16, error, claim) {
	token := ctx.GetHeader("Authorization")

	if len(token) == 0 {
		return false, 401, errors.New("required token"), claim{}
	}

	// The JWTSign is shown for testing purposes
	// The real sign is stored in the .env file
	claims, access, message, err := ProcessToken(token, "secretKeyOnJWT")

	if !access {
		if err != nil {
			fmt.Println("Token error: " + message)
			return false, 401, err, claim{}
		} else {
			fmt.Println("Token error: " + message)
			return false, 401, err, claim{}
		}
	}	

	fmt.Println("Token OK")
	return true, 200, err, *claims

}