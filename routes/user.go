package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)


func signup(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


	err = user.Save()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusCreated, gin.H{"message":"User created successfully"})

}

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


	err = user.ValidateCredentials()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	token ,err := utils.GenerateToken(user.Email,user.ID)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }


	context.JSON(http.StatusOK, gin.H{"message":"User login successfully","token":token})

}