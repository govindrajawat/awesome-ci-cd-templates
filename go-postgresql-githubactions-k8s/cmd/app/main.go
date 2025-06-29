package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/gorilla/mux"
	"go-postgresql-githubactions-k8s/internal/handler"
	"go-postgresql-githubactions-k8s/internal/repository/postgres"
	"go-postgresql-githubactions-k8s/internal/service"
	"go-postgresql-githubactions-k8s/pkg/database"
)

func main() {
	// Initialize database
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repository, service and handler
	messageRepo := postgres.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepo)
	messageHandler := handler.NewMessageHandler(messageService)
	healthHandler := handler.NewHealthHandler(db)

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/api/messages", messageHandler.GetMessages).Methods("GET")
	r.HandleFunc("/api/messages", messageHandler.CreateMessage).Methods("POST")
	r.HandleFunc("/health", healthHandler.HealthCheck).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}