package postgres

import (
	"database/sql"
	"go-postgresql-githubactions-k8s/internal/repository"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) GetAll() ([]repository.Message, error) {
	rows, err := r.db.Query("SELECT id, content FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []repository.Message
	for rows.Next() {
		var m repository.Message
		if err := rows.Scan(&m.ID, &m.Content); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (r *MessageRepository) Create(content string) (repository.Message, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO messages (content) VALUES ($1) RETURNING id",
		content,
	).Scan(&id)

	if err != nil {
		return repository.Message{}, err
	}

	return repository.Message{ID: id, Content: content}, nil
}