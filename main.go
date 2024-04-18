package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InventoryItem defines the structure for an inventory item
type InventoryItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func main() {
	router := gin.Default()

	// Basic route for the homepage
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to SupplySpy!"})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
