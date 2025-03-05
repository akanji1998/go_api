package routes

import (
	"net/http"
	"strconv"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


func registerForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	event,err := models.GetEventById(eventId)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	err = event.Register(userId)

	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusCreated, gin.H{"message":"Registration made"})

}

func cancelRegistration(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	event,err := models.GetEventById(eventId)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	err = event.CancelRegistration(userId)

	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusOK, gin.H{"message":"Registration DELETED"})

}