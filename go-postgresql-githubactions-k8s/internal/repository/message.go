package repository

import "time"

type Message struct {
	ID        int       `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type MessageRepository interface {
	GetAll() ([]Message, error)
	Create(content string) (Message, error)
} 