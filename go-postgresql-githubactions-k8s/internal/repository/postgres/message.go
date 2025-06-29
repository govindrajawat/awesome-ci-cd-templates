package postgres

import (
	"database/sql"
	"go-postgresql-githubactions-k8s/internal/repository"
	"time"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) GetAll() ([]repository.Message, error) {
	rows, err := r.db.Query("SELECT id, content, created_at FROM messages ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []repository.Message
	for rows.Next() {
		var m repository.Message
		if err := rows.Scan(&m.ID, &m.Content, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (r *MessageRepository) Create(content string) (repository.Message, error) {
	var id int
	var createdAt time.Time
	err := r.db.QueryRow(
		"INSERT INTO messages (content) VALUES ($1) RETURNING id, created_at",
		content,
	).Scan(&id, &createdAt)

	if err != nil {
		return repository.Message{}, err
	}

	return repository.Message{ID: id, Content: content, CreatedAt: createdAt}, nil
}