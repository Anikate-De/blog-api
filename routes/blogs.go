package routes

import (
	"log"
	"net/http"
	"strconv"

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

func getBlog(context *gin.Context) {
	id, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted",
		})
		return
	}

	blog, err := models.GetBlogById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Blog does not exist.",
		})
		return
	}

	context.JSON(http.StatusOK, blog)
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
