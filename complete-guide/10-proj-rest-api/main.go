package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oskiegarcia/event-booking/db"
	"github.com/oskiegarcia/event-booking/handlers"
)

func main() {
	db.InitDB()

	server := gin.Default()

	handlers.RegisterRoutes(server)

	server.Run(":8080")
}
