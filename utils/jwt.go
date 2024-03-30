package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	EMAIL_CLAIM   = "X-E"
	USER_ID_CLAIM = "X-U"
	SECRET        = "utc5rpd@dhc7TZV9tmq"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		EMAIL_CLAIM:   email,
		USER_ID_CLAIM: userId,
		"exp":         time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRET))
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	if isTokenValid := parsedToken.Valid; !isTokenValid {
		return 0, errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("claims not ok")
	}

	// email := claims[EMAIL_CLAIM]
	userId := int64(claims[USER_ID_CLAIM].(float64))

	return userId, nil
}
