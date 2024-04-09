package routes

import (
	"net/http"

	"de.anikate/blog-api/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {

	// Home
	engine.GET("/", home)

	allRoutes := engine.Group("/")
	allRoutes.Use(middleware.Auth)

	// Create a new user
	engine.POST("/signup", signup)

	// Login
	engine.POST("/login", login)
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message":    "Welcome to the Go Blog API!",
		"author":     "Anikate De",
		"author_url": "https://github.com/Anikate-De",
		"version":    "0.0.1",
	})
}
