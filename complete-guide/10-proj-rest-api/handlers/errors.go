package handlers

import "github.com/gin-gonic/gin"

func toJSON(err error) gin.H {
	return gin.H{"message": err.Error()}
}
