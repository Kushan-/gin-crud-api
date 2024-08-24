package routes

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"example.com/gin-go-api/models"
	db "example.com/gin-go-api/sql-db"
	"github.com/gin-gonic/gin"
)

func dbConnect(cntx *gin.Context) {
	db.InitDb()
}

func grettings(cntx *gin.Context) {
	cntx.JSON(http.StatusOK, gin.H{
		"txt": "Hello form GIN ===GO!",
	})
}

func getEvents(cntx *gin.Context) {
	events, err := models.GetAllQLEvents()
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

	payload := cntx.Request.Body
	fmt.Println(payload)
	//fmt.Println("Reqs->>>", cntx.Request)

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

func updateEvent(cntx *gin.Context) {
	// defer wg.Done()
	var event models.Event
	idStr := cntx.Param("id")
	id, err := strIdConversion(idStr)

	err = cntx.ShouldBindJSON(&event)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = event.UpdateQLEvent(id)
	if err != nil {
		fmt.Println("updating event err, ", err)
		return
	}

	cntx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(cntx *gin.Context) {
	// defer wg.Done()
	var event models.Event

	idStr := cntx.Param("id")
	id, err := strIdConversion(idStr)
	if err != nil {
		fmt.Println("updating event err, ", err)
		return
	}

	err = event.DeleteQLEvent(id)
	if err != nil {
		fmt.Println("DELETING event err, ", err)
		return
	}

	cntx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func getEvent(cntx *gin.Context) {
	idStr := cntx.Param("id")
	id, err := strIdConversion(idStr)
	if err != nil {
		fmt.Println("updating event err, ", err)
		return
	}

	event, err := models.GetQLEventsById(id)

	if err != nil {
		fmt.Println("GET BY ID event err, ", err)
	}

	cntx.JSON(http.StatusOK, event)

}

func strIdConversion(idStr string) (id int64, err error) {

	id, err = strconv.ParseInt((idStr), 10, 64)

	if err != nil {
		fmt.Println("err while conversion, ", err)
		return
	}
	return id, nil
}
