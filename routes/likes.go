package routes

import (
	"log"
	"net/http"
	"strconv"

	"de.anikate/blog-api/models"
	"github.com/gin-gonic/gin"
)

func likeBlog(context *gin.Context) {
	blogId, err := strconv.ParseInt(context.Params.ByName("bid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted",
		})
		return
	}

	uid := context.GetInt64("uid")

	_, err = models.GetBlogById(blogId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Blog does not exist.",
		})
		return
	}

	like := models.Like{
		AuthorId: uid,
		BlogId:   blogId,
	}

	exists := like.Exists()
	if exists {
		context.JSON(http.StatusOK, gin.H{
			"message": "Blog has already been liked",
			"id":      like.Id,
		})
		return
	}

	err = like.Save()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to like blog.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Blog liked successfully!",
		"id":      like.Id,
	})
}

func unlikeBlog(context *gin.Context) {
	blogId, err := strconv.ParseInt(context.Params.ByName("bid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog ID is ill-formatted",
		})
		return
	}

	uid := context.GetInt64("uid")

	_, err = models.GetBlogById(blogId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Blog does not exist.",
		})
		return
	}

	like := models.Like{
		AuthorId: uid,
		BlogId:   blogId,
	}

	exists := like.Exists()
	if !exists {
		context.JSON(http.StatusOK, gin.H{
			"message": "Blog has already been un-liked",
		})
		return
	}

	err = like.Delete()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to un-like blog.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Blog un-liked successfully!",
	})
}
