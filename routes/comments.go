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
	blogId, err := strconv.ParseInt(context.Params.ByName("bid"), 10, 64)
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

func editComment(context *gin.Context) {
	blogId, err := strconv.ParseInt(context.Params.ByName("bid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted",
		})
		return
	}

	commentId, err := strconv.ParseInt(context.Params.ByName("cid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Comment ID is ill-formatted",
		})
		return
	}

	uid := context.GetInt64("uid")

	comment, err := models.GetCommentById(commentId, blogId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Comment does not exist.",
		})
		return
	}

	if comment.AuthorId != uid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update comment.",
		})
		return
	}

	receivedComment := *comment

	err = context.ShouldBindJSON(&receivedComment)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid POST body.",
		})
		return
	}

	comment.Body = receivedComment.Body

	err = comment.Update()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update comment.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Updated comment successfully!",
	})
}
