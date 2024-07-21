package middlewares

import (
	"net/http"

	"eventbooking.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Cool stuff here
	// Since we are using a pointer for gin Context everywhere we can pass the user ID var to the context itself
	context.Set("userId", userId)
	context.Next()
}
