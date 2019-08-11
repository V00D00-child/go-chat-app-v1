package core

import (
	"fmt"
	"realtime-chat-go-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO: Handle errors in this file and up date token time back to 1 day
var signingKey = "404fbcb51ad427dfd107d406a5175e643a03953fb91a38fe31b57250c4deec4b"

// CreateToken ...
func CreateToken(user models.User) (string, error) {

	expirationTime := time.Now().Add(1 * time.Minute) // unix milliseconds 1440

	claims := Token{
		Username:  user.UserName,
		ID:        user.ID,
		UserEmail: user.UserEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// RenewToken .. Renew token in the last 30 seconds of expiration
func RenewToken(claim *Token) (string, bool) {

	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return "", false
	}
	// renew the time
	claim.ExpiresAt = time.Now().Add(1 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	return tokenString, true

}

// CheckToken ...
func CheckToken(tknStr string) (*Token, error) {
	claims := &Token{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if !tkn.Valid {
		fmt.Println("Token is not valid", tkn.Valid)
		return claims, err
	}

	return claims, nil
}

// Token ...
type Token struct {
	Username  string `json:"username"`
	ID        int    `json:"id"`
	UserEmail string `json:"userEmail"`
	jwt.StandardClaims
}
