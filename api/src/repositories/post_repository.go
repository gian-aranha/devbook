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
func (r Posts) Get(userID uint64) ([]models.Post, error) {
	lines, err := r.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on p.author_id = u.id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?`,
		userID, 
		userID,
	)
	if err != nil {
		return []models.Post{}, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
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

// GetByUserID returns all posts from the user who attends to the received id
func (r Posts) GetByUserID(userID uint64) ([]models.Post, error) {
	lines, err := r.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on p.author_id = u.id
		where p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update alters the post information in the database
func (r Posts) Update(postID uint64, post models.Post) error {
	statement, err := r.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

// Delete removes the post with the received id from the database
func (r Posts) Delete(postID uint64) error {
	statement, err := r.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}