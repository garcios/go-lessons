package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oskiegarcia/event-booking/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.GET("/events/:id/register", listRegisteredUsers)

	// users
	server.POST("signup", signUp)
	server.POST("login", login)
}
