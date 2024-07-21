package middleware

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger writes request informations on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verifies if the user that's making the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidateToken(r); erro != nil {
			responses.Error(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}