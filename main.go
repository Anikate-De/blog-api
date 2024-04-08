package main

import (
	"net/http"

	"de.anikate/blog-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	engine := gin.Default()

	engine.GET("/", home)

	engine.Run(":8080")
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message":    "Welcome to the Go Blog API!",
		"author":     "Anikate De",
		"author_url": "https://github.com/Anikate-De",
		"version":    "0.0.1",
	})
}
