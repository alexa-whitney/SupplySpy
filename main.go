package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Basic route for the homepage
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to SupplySpy!"})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
