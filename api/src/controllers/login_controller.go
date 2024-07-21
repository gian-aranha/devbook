package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login is responsible for authenticating an user in the API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	databaseUser, erro := repository.GetByEmail(user.Email)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(databaseUser.Password, user.Password); erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CreateToken(databaseUser.ID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	
	responses.JSON(w, http.StatusAccepted, token)
}