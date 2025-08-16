package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("supersecret") // later read from config/env

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
