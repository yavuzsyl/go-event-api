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

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return err
	}

	if isTokenValid := parsedToken.Valid; !isTokenValid {
		return errors.New("token is not valid")
	}

	// if claims, ok := token.Claims.(jwt.MapClaims); !ok {
	// 	return errors.New("claims not ok")
	// }

	// email := claims[EMAIL_CLAIM]
	// userId := claims[USER_ID_CLAIM]

	return nil
}
