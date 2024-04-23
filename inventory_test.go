package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// Assuming inventoryFile is a global variable in the main application
var inventoryFile = "inventory.json"

func TestLoadInventory(t *testing.T) {
	testFile := "test_inventory.json"
	defer os.Remove(testFile)

	expected := []InventoryItem{{ID: "1", Name: "Test Item", Description: "A test item", Quantity: 10}}
	data, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal data: %v", err)
	}

	if err := os.WriteFile(testFile, data, 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	oldFile := inventoryFile
	inventoryFile = testFile
	defer func() { inventoryFile = oldFile }()

	loadInventory()
	if !reflect.DeepEqual(inventory, expected) {
		t.Errorf("Expected inventory to be %v, got %v", expected, inventory)
	}
}

func TestSaveInventory(t *testing.T) {
	testFile := "test_inventory.json"
	defer os.Remove(testFile)

	inventory = []InventoryItem{{ID: "2", Name: "Saved Item", Description: "A saved item", Quantity: 5}}
	oldFile := inventoryFile
	inventoryFile = testFile
	defer func() { inventoryFile = oldFile }()

	saveInventory()

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	var savedInventory []InventoryItem
	if err := json.Unmarshal(data, &savedInventory); err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	if !reflect.DeepEqual(savedInventory, inventory) {
		t.Errorf("Expected saved file to contain %v, got %v", inventory, savedInventory)
	}
}

func TestInventoryAPI(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		method     string
		target     string
		body       string
		wantStatus int
	}{
		{"Get Inventory", "GET", "/inventory", "", http.StatusOK},
		{"Add Item", "POST", "/inventory", `{"id":"3","name":"New Item","description":"A new item","quantity":20}`, http.StatusOK},
		{"Delete Item", "DELETE", "/inventory/1", "", http.StatusOK},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.target, bytes.NewBufferString(tt.body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != tt.wantStatus {
			t.Errorf("%s: expected status %d, got %d", tt.name, tt.wantStatus, resp.Code)
		}
	}
}

func BenchmarkInventoryList(b *testing.B) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/inventory", nil)

	for i := 0; i < b.N; i++ {
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
	}
}
