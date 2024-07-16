package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Get returns all users that attends a name or nick filter
func (r Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, erro := r.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, 
		nameOrNick,
	)
	if erro != nil {
		return nil , erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil , erro
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID returns a user that attends to the id received
func (r Users) GetByID(userID uint64) (models.User, error) {
	lines, erro := r.db.Query(
		"select id, name, nick, email, created_at from users where id=?",
		userID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// Update alters the user informations in the database
func (r Users) Update(userID uint64, user models.User) error {
	statement, erro := r.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, userID); erro != nil {
		return erro
	}

	return nil
}

// Delete removes the user with the received id from the database
func (r Users) Delete(userID uint64) error {
	statement, erro := r.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(userID); erro != nil {
		return erro
	}

	return nil
}