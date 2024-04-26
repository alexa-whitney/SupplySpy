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
var (
	inventoryFile = "inventory.json"
	inventory     []InventoryItem
)

// loadInventory loads the inventory from the current inventoryFile
func loadInventory() {
	// Read the inventory file
	data, err := os.ReadFile(inventoryFile)
	// Check if the file does not exist
	if err != nil {
		// Log a message and return
		log.Printf("Could not load inventory from %s, starting with an empty list: %v", inventoryFile, err)
		return
	}
	// Unmarshal the inventory data into the inventory slice
	err = json.Unmarshal(data, &inventory)
	if err != nil {
		log.Fatalf("Error unmarshalling inventory data from %s: %v", inventoryFile, err)
	}
}

// saveInventory saves the inventory to the current inventoryFile
func saveInventory() {
	data, err := json.Marshal(inventory)
	if err != nil {
		log.Fatalf("Error marshalling inventory data: %v", err)
	}
	err = os.WriteFile(inventoryFile, data, 0644)
	if err != nil {
		log.Fatalf("Error writing inventory data to %s: %v", inventoryFile, err)
	}
}

// main is the entry point for the application
func main() {
	// Initialize the router
	router := setupRouter()
	router.Run(":8080") // Start the server on port 8080
}

// setupRouter initializes and returns a new gin router with all the configured routes.
func setupRouter() *gin.Engine {
	// Initialize the router
	router := gin.Default()
	// Serve static files from the static directory
	router.Static("/static", "./static")
	loadInventory() // Load inventory on router setup

	// DEFINE ROUTES //

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
	// Bind the request body to a new InventoryItem
	var newItem InventoryItem
	// Check if the request body can be bound to the new item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if an item with the same ID already exists
	for _, item := range inventory {
		// If an item with the same ID exists, return an error
		if item.ID == newItem.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item with this ID already exists"})
			return
		}
	}
	// Append the new item to the inventory
	inventory = append(inventory, newItem)
	saveInventory()                // Save the updated inventory
	c.JSON(http.StatusOK, newItem) // Return the new item
}

// handleUpdateInventory handles the PUT request to update an existing inventory item
func handleUpdateInventory(c *gin.Context) {
	id := c.Param("id")           // Get the ID from the URL parameter
	var updatedItem InventoryItem // Create a new InventoryItem to hold the updated item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if an item with the same ID already exists
	for i, item := range inventory {
		if item.ID == id {
			inventory[i] = updatedItem
			saveInventory()
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	// If the item was not found, return an error
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// handleDeleteInventory handles the DELETE request to delete an inventory item
func handleDeleteInventory(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL parameter
	// Iterate over the inventory to find the item with the matching ID
	for i, item := range inventory {
		// If the item is found, remove it from the inventory
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			saveInventory()
			c.Status(http.StatusOK)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
