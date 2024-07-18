package main

import (
	"eventbooking.com/rest-api/db"
	"eventbooking.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":3000")
}
