package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := `
    INSERT INTO snippets (title, content, created, expires)
    VALUES (
      ?,
      ?,
      UTC_TIMESTAMP(),
      DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY)
    );
  `
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	s := &Snippet{}
	stmt := `
		SELECT id, title, content, created, expires
		FROM snippets
		WHERE id = ? AND expires > UTC_TIMESTAMP();
	`
	err := m.DB.QueryRow(stmt, id).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.Created,
		&s.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil // TODO: implement this
}
