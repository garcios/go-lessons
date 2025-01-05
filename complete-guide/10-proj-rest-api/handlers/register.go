package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	apperrors "github.com/oskiegarcia/event-booking/app-errors"
	"github.com/oskiegarcia/event-booking/repositories"
	"net/http"
	"strconv"
)

type UsersForEvent struct {
	Event *repositories.Event            `json:"event"`
	Users []*repositories.RegisteredUser `json:"users"`
}

func registerForEvent(c *gin.Context) {
	userID, eventID, err := getUserIDAndEventID(c)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			c.JSON(http.StatusNotFound, toJSON(err))
			return
		}

		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	var registration *repositories.Registration

	registration, err = repositories.FindRegistrationByUserIDAndEventID(userID, eventID)
	if err != nil && !errors.Is(err, apperrors.RegistrationNotFoundError) {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	if registration != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "user already registered for this event"})
		return
	}

	registration = &repositories.Registration{
		EventID: eventID,
		UserID:  userID,
	}

	err = registration.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusCreated, registration)
}

func getUserIDAndEventID(c *gin.Context) (int64, int64, error) {
	userID := c.GetString("userID")

	userIDAsInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	eventID, err := getEventIDFromParam(c)
	if err != nil {
		return 0, 0, err
	}

	_, err = repositories.GetEventByID(eventID)
	if err != nil {
		return 0, 0, err
	}

	return userIDAsInt, eventID, nil
}

func cancelRegistration(c *gin.Context) {
	userID, eventID, err := getUserIDAndEventID(c)
	if err != nil {
		if errors.Is(err, apperrors.EventNotFoundError) {
			c.JSON(http.StatusNotFound, toJSON(err))
			return
		}

		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	registration, err := repositories.FindRegistrationByUserIDAndEventID(userID, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	err = registration.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration cancelled successfully"})
}

func listRegisteredUsers(c *gin.Context) {
	eventID, err := getEventIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
		return
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

	users, err := repositories.GetRegisteredUsersForEventID(eventID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	usersForEvent := UsersForEvent{
		Event: event,
		Users: users,
	}

	c.JSON(http.StatusOK, usersForEvent)
}
