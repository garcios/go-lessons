package middlewares

import (
	"github.com/gin-gonic/gin"
	apperrors "github.com/oskiegarcia/event-booking/app-errors"
	"github.com/oskiegarcia/event-booking/utils"
	"net/http"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": apperrors.NotAuthorizedError.Error()})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": apperrors.NotAuthorizedError.Error()})
		return
	}

	c.Set("userID", userID)

	c.Next()
}
