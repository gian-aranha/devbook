package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post in the application
type Post struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"authorId,omitempty"`
	AuthorNick string `json:"authorNick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreatedAt  time.Time `json:"createAt,omitempty"`
}

// Prepare calls validate and format methods to act uppon the received post
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()

	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title is mandatory and can't be blank")
	}

	if post.Content == "" {
		return errors.New("the content is mandatory and can't be blank")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}