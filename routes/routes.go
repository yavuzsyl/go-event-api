package routes

import (
	"eventapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authorizedEventsGroup := server.Group("/")
	authorizedEventsGroup.Use(middlewares.Authenticate)
	authorizedEventsGroup.GET("/events", getEvents)
	authorizedEventsGroup.GET("/events/:id", getEvent)
	authorizedEventsGroup.POST("/events", createEvent)
	authorizedEventsGroup.PUT("/events/:id", updateEvent)
	authorizedEventsGroup.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
