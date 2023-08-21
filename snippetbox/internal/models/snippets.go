package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int `db:"id"`
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `insert into snippets (title, content, created, expires) 
	values ($1, $2, current_timestamp, current_timestamp + '%d day') returning id;`
	stmt = fmt.Sprintf(stmt, expires)

	var id int
	err := m.DB.QueryRow(stmt, title, content).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting snippet into database: %v", err)
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	s := &Snippet{}

	query := `select id , title, content, created, expires from snippets 
        	  where id = $1`

	err := m.DB.QueryRow(query, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
             WHERE expires > CURRENT_TIMESTAMP
             ORDER BY id DESC
             LIMIT 10;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}
	for rows.Next() {
		s := &Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
