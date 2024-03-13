package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      uint
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

type SnippetModelInterface interface {
	Insert(title, content string, expires uint16) (uint, error)
	Get(id uint) (Snippet, error)
	Latest() ([]Snippet, error)
}

func (m *SnippetModel) Insert(title, content string, expires uint16) (uint, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (m *SnippetModel) Get(id uint) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		}
		return Snippet{}, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	limit := 10
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT ?`

	rows, err := m.DB.Query(stmt, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := make([]Snippet, 0, limit)

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	return snippets, nil
}
