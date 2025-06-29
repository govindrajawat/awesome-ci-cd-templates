package handler

import (
	"encoding/json"
	"net/http"
	
	"go-postgresql-githubactions-k8s/internal/service"
)

type MessageHandler struct {
	service *service.MessageService
}

func NewMessageHandler(s *service.MessageService) *MessageHandler {
	return &MessageHandler{service: s}
}

func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message struct {
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(message.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}