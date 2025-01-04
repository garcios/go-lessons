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

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, toJSON(eventNotFoundError))
		return
	}

	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, toJSON(eventNotFoundError))
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
