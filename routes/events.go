package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusOK,events)
}
func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("event_id"),10,64) 
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
event,err := models.GetEventById(eventId)
if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusOK,event)
	
}


func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("event_id"),10,64) 
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

userId := context.GetInt64("userId")
	event,err := models.GetEventById(eventId)



if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

		if event.UserID != userId {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update event"})
            return
			
		}



	var updatedEvent models.Event


	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	
}


func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("event_id"),10,64) 
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	userId := context.GetInt64("userId")
event,err := models.GetEventById(eventId)

if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	if event.UserID != userId {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to delete event"})
            return
			
		}


	err = event.Delete()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
	
}

func createEvent(context *gin.Context){


	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	context.JSON(http.StatusCreated, gin.H{"message":"Event created", "event":event})

}