package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("secret_key_change_later")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}
