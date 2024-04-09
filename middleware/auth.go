package middleware

import (
	"net/http"
	"strings"

	"de.anikate/blog-api/utils"
	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context) {
	authToken := context.GetHeader("Authorization")
	id, err := utils.ParseJWT(authToken)

	if err != nil {
		if strings.Contains(err.Error(), "token expired") {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid authentication token, expired.")
			return
		}
		context.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid authentication token")
		return
	}

	context.Set("id", id)
	context.Next()
}
