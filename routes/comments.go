package routes

import (
	"log"
	"net/http"
	"strconv"

	"de.anikate/blog-api/models"
	"github.com/gin-gonic/gin"
)

func addComment(context *gin.Context) {

	blogId, err := strconv.ParseInt(context.Params.ByName("bid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted.",
		})
		return
	}

	var comment *models.Comment = &models.Comment{
		AuthorId: context.Value("uid").(int64),
		BlogId:   blogId,
	}

	err = context.ShouldBindJSON(&comment)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body, string field body is required.",
		})
		return
	}

	err = comment.Save()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to post comment.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Added comment successfully!",
		"comment": comment,
	})
}

func getBlogComments(context *gin.Context) {
	blogId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted.",
		})
		return
	}

	comments, err := models.AllComments(blogId)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to retrieve comments.",
		})
		return
	}

	context.JSON(http.StatusOK, comments)
}
