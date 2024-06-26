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
	// Get User by ID
	engine.GET("/users/:uid", getUser)

	// Un-register a user
	authenticated.DELETE("/unregister", unRegister)
	// Update User
	authenticated.PUT("/users/update", updateUser)

	// Get all blogs
	engine.GET("/blogs", getAllBlogs)
	// Get a blog by ID
	engine.GET("/blogs/:bid", getBlog)

	// Create a blog
	authenticated.POST("/blogs", postBlog)
	// Update a blog
	authenticated.PUT("/blogs/:bid", updateBlog)
	// Delete a blog
	authenticated.DELETE("/blogs/:bid", deleteBlog)

	// Get on comments on a blog
	engine.GET("/blogs/:bid/comments", getBlogComments)

	// Comment on a blog
	authenticated.POST("/blogs/:bid/comments", addComment)
	// Update a comment
	authenticated.PUT("/blogs/:bid/comments/:cid", editComment)
	// Delete a comment
	authenticated.DELETE("/blogs/:bid/comments/:cid", deleteComment)

	// Get all likes on a Blog
	engine.GET("/blogs/:bid/likes", getAllLikes)

	// Like a Blog
	authenticated.POST("/blogs/:bid/likes", likeBlog)
	// Unlike a Blog
	authenticated.DELETE("/blogs/:bid/likes", unlikeBlog)

	// Share a blog
	engine.PUT("/blogs/:bid/share", shareBlog)
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message":    "Welcome to the Go Blog API!",
		"author":     "Anikate De",
		"author_url": "https://github.com/Anikate-De",
		"version":    "0.0.1",
	})
}
