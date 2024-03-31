package routes

import (
	"eventapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authorizedEventsGroup := server.Group("/")
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authorizedEventsGroup.Use(middlewares.Authenticate)
	authorizedEventsGroup.POST("/events", createEvent)
	authorizedEventsGroup.PUT("/events/:id", updateEvent)
	authorizedEventsGroup.DELETE("/events/:id", deleteEvent)

	authorizedEventsGroup.POST("/events/:id/register", registerForEvent)
	authorizedEventsGroup.DELETE("/events/:id/cancel-registration", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
