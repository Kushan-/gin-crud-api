package routes

import (
	"example.com/gin-go-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", grettings)
	server.GET("/events", getEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/update/:id", middlewares.Authenticate, updateEvent)
	authenticated.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/update/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	server.GET("/events/:id", getEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
