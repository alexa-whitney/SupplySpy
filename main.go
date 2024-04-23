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
	router := gin.Default()

	// Serve static files like CSS, JavaScript, or images
	router.Static("/static", "./static")

	// Load existing inventory, if any
	loadInventory()

	// Route for home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Welcome to Awesome New Startup, Inc.",
		})
	})

	// Route to list all items with a template
	router.GET("/inventory", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inventory.html", gin.H{
			"Inventory": inventory,
			"title":     "Inventory List",
		})
	})

	// Route to serve the add item form with template
	router.GET("/add-item", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_item.html", gin.H{
			"title": "Add New Inventory Item",
		})
	})

	// Route to add an item via POST request from the form
	router.POST("/inventory", func(c *gin.Context) {
		// Bind the JSON data to the InventoryItem struct
		var newItem InventoryItem
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if item with the same ID already exists
		for _, item := range inventory {
			// If item with the same ID already exists, return an error
			if item.ID == newItem.ID {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Item with this ID already exists"})
				return
			}
		}

		// Append the new item to the inventory
		inventory = append(inventory, newItem)
		// Save the inventory to a JSON file
		saveInventory()
		// Return the new item as JSON
		c.JSON(http.StatusOK, newItem)
	})

	// Route to get a single item by ID
	router.PUT("/inventory/:id", func(c *gin.Context) {
		// Get the ID from the URL
		id := c.Param("id")
		// Bind the JSON data to the InventoryItem struct
		var updatedItem InventoryItem
		if err := c.ShouldBindJSON(&updatedItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find and update the item
		for i, item := range inventory {
			// If item with the same ID already exists, return an error
			if item.ID == id {
				// Update the item
				inventory[i] = updatedItem
				saveInventory()
				c.JSON(http.StatusOK, updatedItem)
				return
			}
		}

		// If item with the ID was not found, return an error
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Route to update an item via PUT request from the edit form
	router.POST("/inventory/:id/update", func(c *gin.Context) {
		id := c.Param("id")
		var updatedItem InventoryItem
		if err := c.ShouldBindJSON(&updatedItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Find and update the item
		for i, item := range inventory {
			if item.ID == id {
				inventory[i] = updatedItem
				saveInventory()
				c.Redirect(http.StatusFound, "/inventory")
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Route to delete an item by ID
	router.DELETE("/inventory/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Find and delete the item
		for i, item := range inventory {
			if item.ID == id {
				inventory = append(inventory[:i], inventory[i+1:]...)
				saveInventory()
				c.Status(http.StatusOK)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Load HTML files from the templates directory
	router.LoadHTMLGlob("templates/*")

	// Start the server on port 8080
	router.Run(":8080")
}
