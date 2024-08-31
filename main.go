package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.com/db"
	"restapi.com/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
