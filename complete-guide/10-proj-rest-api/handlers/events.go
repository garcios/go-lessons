package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oskiegarcia/event-booking/models"
	"net/http"
	"strconv"
)

var (
	eventNotFoundError = errors.New("event not found")
)

func toJSON(err error) gin.H {
	return gin.H{"message": err.Error()}
}

func createEvent(c *gin.Context) {
	var event models.Event

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

	events, err := models.GetAllEvents()
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

	event, err := retrieveEventOrError(c, eventID)
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

	_, err = retrieveEventOrError(c, eventID)
	if err != nil {
		return
	}

	var updatedEvent models.Event
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

func retrieveEventOrError(c *gin.Context, eventID int64) (*models.Event, error) {
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return nil, err
	}

	if event == nil {
		c.JSON(http.StatusNotFound, toJSON(eventNotFoundError))
		return nil, eventNotFoundError
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

	event, err := retrieveEventOrError(c, eventID)
	if err != nil {
		// No need to set error in the context here because retrieveEventOrError has already done so.
		return
	}

	event.ID = eventID

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
