package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine){

	server.GET("/events",getEvents)
	server.GET("/events/:event_id",getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:event_id",updateEvent)
	authenticated.DELETE("/events/:event_id",deleteEvent)
	authenticated.POST("/event/:id/register",registerForEvent)
	authenticated.DELETE("/event/:id/register",cancelRegistration)

	server.POST("/signup",signup)
	server.POST("/login",login)
	
	
}