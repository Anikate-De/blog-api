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
