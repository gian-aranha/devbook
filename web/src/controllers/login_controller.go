package controllers

import (
	"net/http"
	"web/src/utils"
)

// LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}