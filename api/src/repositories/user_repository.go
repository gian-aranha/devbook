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
	statement, err := r.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(insertedID), nil
}

// Get returns all users that attends a name or nick filter
func (r Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := r.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, 
		nameOrNick,
	)
	if err != nil {
		return nil , err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil , err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID returns a user that attends to the id received
func (r Users) GetByID(userID uint64) (models.User, error) {
	lines, err := r.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// GetByEmail returns a user id and hashed password based on the received email
func (r Users) GetByEmail(userEmail string) (models.User, error) {
	line, err := r.db.Query("select id, password from users where email = ?", userEmail)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update alters the user informations in the database
func (r Users) Update(userID uint64, user models.User) error {
	statement, err := r.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

// Delete removes the user with the received id from the database
func (r Users) Delete(userID uint64) error {
	statement, err := r.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// Follow add the userID and the follower ID to the followers table
func (r Users) Follow(userID, followerID uint64) error {
	statement, err := r.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow removes the userID and the followerID to the followers table
func (r Users) Unfollow(userID, followerID uint64) error {
	statement, err := r.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// GetFollowers returns all followers from the given user
func (r Users) GetFollowers(userID uint64) ([]models.User, error) {
	lines, err := r.db.Query(
		`select u.id, u.name, u.nick, u.email, u.created_at from users u 
		inner join followers f on u.id = f.follower_id where f.user_id = ?`,
		 userID,
	)
	if err != nil {
		return []models.User{}, err
	}
	defer lines.Close()

	var users []models.User 

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name, 
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetFollowing returns all users the given user follows
func (r Users) GetFollowing(userID uint64) ([]models.User, error) {
	lines, err := r.db.Query(
		`select u.id, u.name, u.nick, u.email, u.created_at from users u 
		inner join followers f on u.id = f.user_id where f.follower_id = ?`,
		userID,
	)
	if err != nil {
		return []models.User{}, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User 

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}