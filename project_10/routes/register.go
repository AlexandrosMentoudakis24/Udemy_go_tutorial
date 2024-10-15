package routes

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event!"})

		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to register event!"},
		)

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event registered!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

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

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, 
			gin.H{"message": "Failed to cancel event registration!"},
		)

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event registration cancelled!"})
}

