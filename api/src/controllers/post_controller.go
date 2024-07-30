package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CreatePost creates a post in the database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	var user models.User
	userRepository := repositories.NewUsersRepository(db)
	user, err = userRepository.GetByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	post.AuthorID = userID
	post.AuthorNick = user.Nick

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts gets all posts in the database
func GetPosts(w http.ResponseWriter, r *http.Request) {

}

// GetPost gets a post with specified id in the database
func GetPost(w http.ResponseWriter, r *http.Request) {

}

// UpdatePost updates a post with specified id in the database
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// DeletePost deletes a post with specified id in the database
func DeletePost(w http.ResponseWriter, r *http.Request) {

}