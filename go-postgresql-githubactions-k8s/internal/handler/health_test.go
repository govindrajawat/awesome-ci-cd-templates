package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"database/sql"
)

func TestHealthHandler_HealthCheck(t *testing.T) {
	// Create a mock database connection
	// In a real test, you might want to use a test database
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		// Skip test if database is not available
		t.Skip("Database not available for testing")
	}
	defer db.Close()
	
	handler := NewHealthHandler(db)
	
	// Create a test request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	// Create a response recorder
	rr := httptest.NewRecorder()
	
	// Call the handler
	handler.HealthCheck(rr, req)
	
	// Check the status code
	if status := rr.Code; status != http.StatusOK && status != http.StatusServiceUnavailable {
		t.Errorf("handler returned unexpected status code: got %v", status)
	}
	
	// Check that we got a JSON response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json", contentType)
	}
} 