package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.com/utils"
)

func Authenticate(context *gin.Context) {
	token := utils.ExtractToken(context.Request)
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
