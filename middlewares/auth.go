package middlewares

import (
	"eventapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {

		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is invalid"})
	}

	context.Set("userId", userId)
	context.Next()
}
