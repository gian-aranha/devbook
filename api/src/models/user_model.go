package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("the name is mandatory and can't be blank")
	}

	if user.Nick == "" {
		return errors.New("the nick is mandatory and can't be blank")
	}

	if user.Email == "" {
		return errors.New("the email is mandatory and can't be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("the inserted email is invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("the password is mandatory and can't be blank")
	}
	
	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		hashPassword, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(hashPassword)
	}

	return nil
}