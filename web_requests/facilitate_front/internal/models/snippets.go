package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet struct for holding snippet data
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Wrap a sql.DB connection in SnippetModel struct
type SnippetModel struct {
	DB *sql.DB
}

// insert new snippet into database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// execute statement
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	// return ID of new snippet
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

// GET snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// use QueryRow() to run statement and get back one row
	row := m.DB.QueryRow(stmt, id)

	// initialize pointer to new Snippet struct
	s := &Snippet{}

	// row.Scan() to fill in s (pointer to Snippet struct)
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// use Query() to return 1 or more rows
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append to slice of snippets
		snippets = append(snippets, s)
	}

	// check for errors during the iterations
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// return snippets slice
	return snippets, nil
}
