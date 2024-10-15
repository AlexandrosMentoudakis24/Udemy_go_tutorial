package routes

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Could not fetch events!"},
		)

		return
	}

	context.JSON(http.StatusOK, events)
}

func getSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(
		context.Param("id"),
		10,
		64,
	)

	if err != nil {
		context.JSON(
			http.StatusBadRequest, 
			gin.H{"message": "Failed to parse event id!"},
		)

		return
	}

	event, err := models.GetSingleEvent(eventId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to find event!"},
		)

		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse data!"})

		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to create new event!"},
		)

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(
		context.Param("id"),
		10,
		64,
	)

	if err != nil {
		context.JSON(
			http.StatusBadRequest, 
			gin.H{"message": "Failed to parse event id!"},
		)

		return
	}

	event, err := models.GetSingleEvent(eventId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to fetch event!"},
		)

		return

	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Unauthorized user to update event!"},
		)
	
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to parse data!"},
		)

		return

	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to update event!"},
		)

		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(
		context.Param("id"),
		10,
		64,
	)

	if err != nil {
		context.JSON(
			http.StatusBadRequest, 
			gin.H{"message": "Failed to parse event id!"},
		)

		return
	}

	event, err := models.GetSingleEvent(eventId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to find event!"},
		)

		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Unauthorized user to delete event!"},
		)
	
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to delete event!"},
		)

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
