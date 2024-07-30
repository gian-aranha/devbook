package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a post repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create inserts a post in the database
func (r Posts) Create(post models.Post) (uint64, error) {
	statement, err := r.db.Prepare("insert into posts (title, content, author_id, author_nick) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID, post.AuthorNick)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}