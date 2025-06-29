package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Database  string    `json:"database"`
	Uptime    int64     `json:"uptime"`
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	
	// Check database connection
	var dbStatus string
	if err := h.db.Ping(); err != nil {
		dbStatus = "disconnected"
	} else {
		dbStatus = "connected"
	}
	
	health := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Database:  dbStatus,
		Uptime:    time.Since(startTime).Milliseconds(),
	}
	
	if dbStatus == "disconnected" {
		health.Status = "unhealthy"
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
} 