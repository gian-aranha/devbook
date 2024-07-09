package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser creates a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body) 
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewUsersRepository(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))
}

// GetUsers gets all users in the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users..."))
}

// GetUser gets a user with specified id in the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user..."))
}

// UpdateUser updates a user with specified id in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user..."))
}

// DeleteUser deletes a user with specified id in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}