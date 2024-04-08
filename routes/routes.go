package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {

	// Home
	engine.GET("/", home)

	// Create a new user
	engine.POST("/signup", signup)

}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message":    "Welcome to the Go Blog API!",
		"author":     "Anikate De",
		"author_url": "https://github.com/Anikate-De",
		"version":    "0.0.1",
	})
}
