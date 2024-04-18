package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// InventoryItem defines the structure for an inventory item
type InventoryItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

// inventory is a slice of InventoryItem
var inventory []InventoryItem

// func loadInventory loads the inventory from a JSON file
func loadInventory() {
	data, err := os.ReadFile("inventory.json")
	if err != nil {
		log.Println("Could not load existing inventory, starting with an empty list:", err)
		return
	}
	err = json.Unmarshal(data, &inventory)
	if err != nil {
		log.Fatalf("Error unmarshalling inventory data: %v", err)
	}
}

func saveInventory() {
	data, err := json.Marshal(inventory)
	if err != nil {
		log.Fatalf("Error marshalling inventory data: %v", err)
	}
	err = os.WriteFile("inventory.json", data, 0644)
	if err != nil {
		log.Fatalf("Error writing inventory data: %v", err)
	}
}

func main() {
	router := gin.Default()

	// Load existing inventory, if any
	loadInventory()

	// Basic route for the homepage
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to SupplySpy!"})
	})

	// Route to list all items
	router.GET("/inventory", func(c *gin.Context) {
		c.JSON(http.StatusOK, inventory)
	})

	// Route to add an item
	router.POST("/inventory", func(c *gin.Context) {
		var newItem InventoryItem
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		inventory = append(inventory, newItem)
		// Save the new inventory list with the new item
		saveInventory()
		c.JSON(http.StatusOK, newItem)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
