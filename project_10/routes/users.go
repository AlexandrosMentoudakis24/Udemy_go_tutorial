package routes

import (
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse data!"})

		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to create new user!"},
		)

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to create new user!"},
		)

		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(
			http.StatusUnauthorized, 
			gin.H{"message": err.Error()},
		)

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to authenticate user!"},
		)

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in!", "token": token})
}
