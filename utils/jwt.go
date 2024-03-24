package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "utc5rpd@dhc7TZV9tmq"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"X-E": email,
		"X-U": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}
