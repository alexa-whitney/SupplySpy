package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

// Defines the test inventory file globally for the test file
const testInventoryFile = "inventory_test.json"

// Setup a consistent test environment before each test
func setup() *gin.Engine {
	inventoryFile = testInventoryFile
	inventory = []InventoryItem{} // Reset the inventory

	// Initialize inventory with default data
	defaultInventory := []InventoryItem{
		{ID: "1", Name: "Test Item 1", Description: "A description for Test Item 1", Quantity: 10},
		{ID: "2", Name: "Test Item 2", Description: "A description for Test Item 2", Quantity: 20},
	}

	// Marshal the default inventory into JSON
	data, err := json.Marshal(defaultInventory)
	if err != nil {
		log.Fatalf("Failed to marshal default inventory: %v", err)
	}

	// Write the default data to the test inventory file
	err = os.WriteFile(inventoryFile, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write default inventory to test file: %v", err)
	}

	// Load the initial test inventory
	loadInventory()

	return setupRouter()
}

// Teardown after each test to clean up
func teardown() {
	os.Remove(testInventoryFile)
}

// TestLoadInventory tests if the inventory is loaded correctly from a file
func TestLoadInventory(t *testing.T) {
	// Setup
	setup()
	defer teardown()

	// Test loading inventory from the test file
	expected := []InventoryItem{{ID: "1", Name: "Test Item", Description: "A test item", Quantity: 10}}
	data, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal data: %v", err)
	}

	// Write the test data to the test inventory file
	if err := os.WriteFile(testInventoryFile, data, 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Load the inventory from the test file
	loadInventory()

	// Compare the loaded inventory with the expected inventory
	if !reflect.DeepEqual(inventory, expected) {
		t.Errorf("Expected inventory to be %v, got %v", expected, inventory)
	}
}

// TestSaveInventory tests if the inventory is saved successfully to a file
func TestSaveInventory(t *testing.T) {
	// Setup
	setup()
	defer teardown()

	// Test saving inventory to the test file
	inventory = []InventoryItem{{ID: "2", Name: "Saved Item", Description: "A saved item", Quantity: 5}}
	saveInventory()

	// Read the saved file and compare the contents
	data, err := os.ReadFile(testInventoryFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Unmarshal the data into a slice of InventoryItem
	var savedInventory []InventoryItem
	if err := json.Unmarshal(data, &savedInventory); err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	// Compare the saved inventory with the expected inventory
	if !reflect.DeepEqual(savedInventory, inventory) {
		t.Errorf("Expected saved file to contain %v, got %v", inventory, savedInventory)
	}
}

// TestInventoryAPI tests the API endpoints for the inventory (GET, POST, and DELETE requests)
func TestInventoryAPI(t *testing.T) {
	// Setup
	router := setup()
	defer teardown()

	// Pre-populate the inventory for consistent API testing
	inventory = []InventoryItem{
		{ID: "1", Name: "Test Item", Description: "A test item", Quantity: 10},
		{ID: "3", Name: "New Item", Description: "A new item", Quantity: 20},
	}
	saveInventory()

	// Define test cases for the API endpoints
	tests := []struct {
		name       string
		method     string
		target     string
		body       string
		wantStatus int
	}{
		{"Get Inventory", "GET", "/inventory", "", http.StatusOK},
		{"Add Item", "POST", "/inventory", `{"id":"4","name":"Added Item","description":"An added item","quantity":15}`, http.StatusOK},
		{"Delete Item", "DELETE", "/inventory/1", "", http.StatusOK},
	}

	// Run the tests
	for _, tt := range tests {
		// Create a new HTTP request for each test case
		req, err := http.NewRequest(tt.method, tt.target, bytes.NewBufferString(tt.body))
		if err != nil {
			t.Fatalf("Failed to create HTTP request: %v", err)
		}
		// Set the request content type to JSON
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to capture the response
		resp := httptest.NewRecorder()
		// Serve the HTTP request to the router
		router.ServeHTTP(resp, req)

		// Check the status code of the response
		if resp.Code != tt.wantStatus {
			t.Errorf("%s: expected status %d, got %d", tt.name, tt.wantStatus, resp.Code)
		}
	}
}

// Benchmark test to measure the performance of loading the inventory
func BenchmarkInventoryList(b *testing.B) {
	router := setup()
	defer teardown() // Ensure clean up after benchmarking

	// Pre-populate the inventory with some items for the benchmark.
	inventory = []InventoryItem{
		{ID: "1", Name: "Test Item 1", Description: "A description for Test Item 1", Quantity: 10},
		{ID: "2", Name: "Test Item 2", Description: "A description for Test Item 2", Quantity: 20},
		{ID: "3", Name: "Test Item 3", Description: "A description for Test Item 3", Quantity: 30},
	}
	saveInventory() // Save these items to the test inventory JSON file

	// Run the benchmark
	b.ResetTimer() // Start timing after setup
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/inventory", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Optionally check the status code to ensure correct handling
		if resp.Code != http.StatusOK {
			b.Errorf("Expected status code 200, got %d", resp.Code)
		}
	}
}
