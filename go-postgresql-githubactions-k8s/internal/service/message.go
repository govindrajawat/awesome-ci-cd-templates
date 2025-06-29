package service

import (
	"go-postgresql-githubactions-k8s/internal/repository"
)

type MessageService struct {
	repo repository.MessageRepository
}

func NewMessageService(r repository.MessageRepository) *MessageService {
	return &MessageService{repo: r}
}

func (s *MessageService) GetAll() ([]repository.Message, error) {
	return s.repo.GetAll()
}

func (s *MessageService) Create(content string) (repository.Message, error) {
	messages, err := s.repo.GetAll()
	if err != nil {
		return repository.Message{}, err
	}
	
	if len(messages) == 0 {
		if err := s.repo.Create("Hello, World!"); err != nil {
			return repository.Message{}, err
		}
	}
	
	return s.repo.Create(content)
}