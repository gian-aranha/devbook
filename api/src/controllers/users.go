package controllers

import "net/http"

// CreateUser creates a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
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