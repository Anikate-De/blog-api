package routes

import (
	"log"
	"net/http"

	"de.anikate/blog-api/models"
	"github.com/gin-gonic/gin"
)

func getAllBlogs(context *gin.Context) {
	blogs, err := models.AllBlogs()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve blogs.",
		})
		return
	}

	context.JSON(http.StatusOK, blogs)
}

func postBlog(context *gin.Context) {
	var blog *models.Blog = &models.Blog{
		AuthorId: context.Value("uid").(int64),
	}

	err := context.ShouldBindJSON(&blog)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body, string fields title, content, and integer author_id are required.",
		})
		return
	}

	err = blog.Save()
	if err != nil {
		log.Println(blog.AuthorId)
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to post blog",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Blog posted successfully!",
		"id":      blog.Id,
	})
}
