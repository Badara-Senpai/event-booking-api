package routes

import (
	"eventbooking.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)

	// two ways to use the auth middleware to protect some routes
	// 1 - Add the middleware to the routes directly
	// server.POST("/events", middlewares.Authenticate, createEvent)

	// 2 - Group some routes and add the common middleware once
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/events/:id", getEvent)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/signin", signin)
}
