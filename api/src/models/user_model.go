package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user in the application
type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare calls validate and format methods to act uppon the received user
func (user *User) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.format()

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("the name is mandatory and can't be blank")
	}

	if user.Nick == "" {
		return errors.New("the nick is mandatory and can't be blank")
	}

	if user.Email == "" {
		return errors.New("the email is mandatory and can't be blank")
	}

	if user.Password == "" {
		return errors.New("the password is mandatory and can't be blank")
	}
	
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}