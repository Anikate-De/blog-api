package routes

import (
	"log"
	"net/http"
	"strings"

	"de.anikate/blog-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user *models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body, must contain string fields name, email and password.",
		})
		return
	}

	err = user.Save()
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to add user to the database. An account with the email already exists.",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to add user to the database.",
		})
		return
	}

}
