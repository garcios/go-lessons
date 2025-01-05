package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oskiegarcia/event-booking/repositories"
	"github.com/oskiegarcia/event-booking/utils"
	"net/http"
	"strconv"
)

func signUp(c *gin.Context) {
	var user repositories.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, toJSON(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(c *gin.Context) {
	var user repositories.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, toJSON(err))
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, toJSON(err))
		return
	}

	token, err := utils.GenerateToken(user.Email, strconv.Itoa(int(user.ID)))
	if err != nil {
		c.JSON(http.StatusUnauthorized, toJSON(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
