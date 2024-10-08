package routes

import (
	"eventapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	event, shouldReturn := fetchEventByRequest(context)
	if shouldReturn {
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {
	event, shouldReturn := fetchEventByRequest(context)
	if shouldReturn {
		return
	}

	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event update request"})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	event, shouldReturn := fetchEventByRequest(context)
	if shouldReturn {
		return
	}

	err := event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}

// COMMON LOGIC
func fetchEventByRequest(context *gin.Context) (event *models.Event, shouldReturn bool) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return nil, true
	}

	event, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find event with given id"})
		return nil, true
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return nil, true
	}

	return event, false
}
