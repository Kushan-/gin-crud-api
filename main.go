package main

import (
	"example.com/gin-go-api/routes"
	db "example.com/gin-go-api/sql-db"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

func main() {
	server := gin.Default() // set up engine aka http server
	routes.RegisterRoutes(server)
	db.InitDb()
	if err := server.Run(":7156"); err != nil {
		panic("Failed to start the server: " + err.Error())
	}

}
