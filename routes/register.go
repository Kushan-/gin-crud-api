package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", grettings)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.PUT("/update/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.GET("/events/:id", getEvent)

}
