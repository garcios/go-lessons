package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	apperrors "github.com/oskiegarcia/event-booking/app-errors"
	"github.com/oskiegarcia/event-booking/repositories"
	"net/http"
	"strconv"
)

func createEvent(c *gin.Context) {
	var event repositories.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
		return
	}

	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, toJSON(apperrors.NotAuthorizedError))
	}

	parsedInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	event.UserID = parsedInt

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusCreated, event)
}

func getEvents(c *gin.Context) {

	events, err := repositories.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {

	eventID, err := getEventIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	event, err := repositories.GetEventByID(eventID)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			c.JSON(http.StatusNotFound, toJSON(err))
		}

		c.JSON(http.StatusInternalServerError, toJSON(err))

		return
	}

	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {

	eventID, err := getEventIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
		return
	}

	_, err = repositories.GetEventByID(eventID)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			c.JSON(http.StatusNotFound, toJSON(err))
			return
		}

		c.JSON(http.StatusInternalServerError, toJSON(err))

		return
	}

	var updatedEvent repositories.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
		return
	}

	updatedEvent.ID = eventID
	updatedEvent.UserID = 1

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func findEvent(eventID int64) (*repositories.Event, error) {
	event, err := repositories.GetEventByID(eventID)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			return nil, err
		}

		return nil, err
	}

	return event, nil
}

func getEventIDFromParam(c *gin.Context) (int64, error) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	return eventID, err
}

func deleteEvent(c *gin.Context) {
	eventID, err := getEventIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	event, err := repositories.GetEventByID(eventID)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			c.JSON(http.StatusNotFound, toJSON(err))
			return
		}

		c.JSON(http.StatusInternalServerError, toJSON(err))

		return
	}

	event.ID = eventID

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
