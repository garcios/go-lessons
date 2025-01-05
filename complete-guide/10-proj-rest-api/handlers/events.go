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

	event.UserID = 1

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

	eventID, err := getEventID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	event, err := findEvent(c, eventID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {

	eventID, err := getEventID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	_, err = findEvent(c, eventID)
	if err != nil {
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

func findEvent(c *gin.Context, eventID int64) (*repositories.Event, error) {
	event, err := repositories.GetEventByID(eventID)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			c.JSON(http.StatusNotFound, toJSON(err))
			return nil, err
		}

		c.JSON(http.StatusInternalServerError, toJSON(err))
		return nil, err
	}

	return event, nil
}

func getEventID(c *gin.Context) (int64, error) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	return eventID, err
}

func deleteEvent(c *gin.Context) {
	eventID, err := getEventID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	event, err := findEvent(c, eventID)
	if err != nil {
		// No need to set error in the context here because findEvent has already done so.
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
