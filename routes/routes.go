package routes

import (
	"net/http"

	"de.anikate/blog-api/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {

	// Home
	engine.GET("/", home)

	authenticated := engine.Group("/")
	authenticated.Use(middleware.Auth)

	// Create a new user
	engine.POST("/signup", signup)
	// Login
	engine.POST("/login", login)

	// Get all blogs
	engine.GET("/blogs", getAllBlogs)
	// Get a blog by ID
	engine.GET("/blogs/:id", getBlog)

	// Create a blog
	authenticated.POST("/blogs", postBlog)
	// Update a blog
	authenticated.PUT("/blogs/:id", updateBlog)
	// Delete a blog
	authenticated.DELETE("/blogs/:id", deleteBlog)
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message":    "Welcome to the Go Blog API!",
		"author":     "Anikate De",
		"author_url": "https://github.com/Anikate-De",
		"version":    "0.0.1",
	})
}
