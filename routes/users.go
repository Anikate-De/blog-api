package routes

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"de.anikate/blog-api/models"
	"de.anikate/blog-api/utils"
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

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!",
	})
}

func login(context *gin.Context) {
	var user *models.User = &models.User{
		Name: "_",
	}

	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body, must contain string fields email and password.",
		})
		return
	}

	err = user.Authenticate()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid login credentials.",
		})
		return
	}

	authToken, err := utils.GenerateJWT(user.Email, user.Uid)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to generate token.",
		})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Login successful!",
		"token":   authToken,
	})
}

func unRegister(context *gin.Context) {
	uid := context.GetInt64("uid")

	user := models.User{
		Uid: uid,
	}
	err := user.Delete()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "Unable to find user, please refresh token.",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to delete user.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Deleted user successfully!",
	})
}

func getUser(context *gin.Context) {
	uid, err := strconv.ParseInt(context.Params.ByName("uid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "User ID is ill-formatted",
		})
		return
	}

	user, err := models.GetUserByID(uid)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exist.",
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func updateUser(context *gin.Context) {
	uid := context.GetInt64("uid")
	user, err := models.GetUserByID(uid)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to find user, please refresh token.",
		})
		return
	}
	user.Password = "_"

	err = context.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body.",
		})
		return
	}

	user.Uid = uid
	err = user.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Updated user successfully!",
	})
}
