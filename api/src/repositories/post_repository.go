package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
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
	statement, err := r.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// Get returns all posts of users that the logged user follows
func (r Posts) Get() ([]models.Post, error) {
	return []models.Post{}, errors.New("error")
}

// GetByID returns a post that attends to the received id 
func (r Posts) GetByID(postID uint64) (models.Post, error) {
	line, err := r.db.Query(`
		select p.*, u.nick from posts p 
		inner join users u on u.id = p.author_id where p.id = ?`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post

	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title, 
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}