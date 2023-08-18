package models

import (
	"database/sql"
	"fmt"
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

func (m *SnippetModel) GetId(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() (*[]Snippet, error) {
	return nil, nil
}
