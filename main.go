package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/gin-go-api/models"
	db "example.com/gin-go-api/sql-db"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

func main() {
	server := gin.Default() // set up engine aka http server

	server.GET("/", grettings)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/dbConn", dbConnect)

	if err := server.Run(":7156"); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}

func dbConnect(cntx *gin.Context) {
	db.InitDb()
}

func grettings(cntx *gin.Context) {
	cntx.JSON(http.StatusOK, gin.H{
		"txt": "Hello form GIN ===GO!",
	})
}
func getEvents(cntx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		cntx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Getting Events error, couldn't fetch it",
			"data": err,
		})
		return
	}
	cntx.JSON(http.StatusOK, gin.H{
		"msg":  "Getting Events",
		"data": events,
	}) // returning response in json format
}
func createEvent(cntx *gin.Context) {

	bodyBytes, err := io.ReadAll(cntx.Request.Body)

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Log the raw request body
	log.Println("Request Body:", string(bodyBytes))

	// Restore the request body so it can be read again later
	cntx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var event models.Event
	fmt.Println(event)
	err = cntx.ShouldBindJSON(&event)

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "Couldn't parse the payload",
			"data": err,
		})
		return
	}

	event.ID = 1
	event.UserId = 1

	fmt.Println(event.ID, event.UserId)

	payload := cntx.Params
	fmt.Println(payload)
	fmt.Println("Reqs->>>", cntx.Request)

	err = event.SaveToQL()

	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "Couldn't fetch events",
			"data": err,
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"msg":  "Event created",
		"data": event,
	})
}
