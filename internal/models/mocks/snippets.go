package mocks

import (
	"lets-go-snippetbox/internal/models"
	"time"
)

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content string, expires uint16) (uint, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id uint) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil

	default:
		return models.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}
