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

// loadInventory loads the inventory from a JSON file
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
	router := setupRouter()
	router.Run(":8080") // Start the server on port 8080
}

// setupRouter initializes and returns a new gin router with all the configured routes.
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")
	loadInventory() // Load inventory on router setup

	// Define routes

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{"title": "Welcome to Awesome New Startup, Inc."})
	})

	// Inventory route
	router.GET("/inventory", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inventory.html", gin.H{"Inventory": inventory, "title": "Inventory List"})
	})

	// Add item route
	router.GET("/add-item", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_item.html", gin.H{"title": "Add New Inventory Item"})
	})

	// API routes
	router.POST("/inventory", handleAddInventory)
	router.PUT("/inventory/:id", handleUpdateInventory)
	router.POST("/inventory/:id/update", handleUpdateInventory)
	router.DELETE("/inventory/:id", handleDeleteInventory)

	// Load templates
	router.LoadHTMLGlob("templates/*")
	// Return the router
	return router
}

// handleAddInventory handles the POST request to add a new inventory item
func handleAddInventory(c *gin.Context) {
	var newItem InventoryItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, item := range inventory {
		if item.ID == newItem.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item with this ID already exists"})
			return
		}
	}
	inventory = append(inventory, newItem)
	saveInventory()
	c.JSON(http.StatusOK, newItem)
}

// handleUpdateInventory handles the PUT request to update an existing inventory item
func handleUpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var updatedItem InventoryItem
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, item := range inventory {
		if item.ID == id {
			inventory[i] = updatedItem
			saveInventory()
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// handleDeleteInventory handles the DELETE request to delete an inventory item
func handleDeleteInventory(c *gin.Context) {
	id := c.Param("id")
	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			saveInventory()
			c.Status(http.StatusOK)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
