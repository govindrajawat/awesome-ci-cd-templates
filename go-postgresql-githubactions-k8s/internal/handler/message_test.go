package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"go-postgresql-githubactions-k8s/internal/service"
)

func TestMessageHandler_GetMessages(t *testing.T) {
	// Create a mock service
	mockService := &service.MessageService{}
	
	handler := NewMessageHandler(mockService)
	
	// Create a test request
	req, err := http.NewRequest("GET", "/api/messages", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	// Create a response recorder
	rr := httptest.NewRecorder()
	
	// Call the handler
	handler.GetMessages(rr, req)
	
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMessageHandler_CreateMessage(t *testing.T) {
	// Create a mock service
	mockService := &service.MessageService{}
	
	handler := NewMessageHandler(mockService)
	
	// Create test data
	message := struct {
		Content string `json:"content"`
	}{
		Content: "Test message",
	}
	
	jsonData, _ := json.Marshal(message)
	
	// Create a test request
	req, err := http.NewRequest("POST", "/api/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Create a response recorder
	rr := httptest.NewRecorder()
	
	// Call the handler
	handler.CreateMessage(rr, req)
	
	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
} 