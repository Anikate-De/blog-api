package routes

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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

func updateBlog(context *gin.Context) {
	id, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted",
		})
		return
	}

	uid := context.GetInt64("uid")

	blog, err := models.GetBlogById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Blog does not exist.",
		})
		return
	}

	if blog.AuthorId != uid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update blog post.",
		})
		return
	}

	receivedBlog := *blog

	err = context.ShouldBindJSON(&receivedBlog)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid POST body.",
		})
		return
	}

	blog.Title = receivedBlog.Title
	blog.Content = receivedBlog.Content
	err = blog.Update()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update blog.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Updated blog successfully!",
	})
}

func deleteBlog(context *gin.Context) {
	id, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted",
		})
		return
	}

	uid := context.GetInt64("uid")

	blog := models.Blog{
		Id:       id,
		AuthorId: uid,
	}
	err = blog.Delete()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "Unable to delete blog, could not find blog with given id and author details.",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to delete blog.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Deleted blog successfully!",
	})
}
