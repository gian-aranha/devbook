package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository creates a user repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts a user in the database
func (r Users) Create(user models.User) (uint64, error) {
	statement, erro := r.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	insertedID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(insertedID), nil
}