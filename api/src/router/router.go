package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with all routs configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.SetRoutes(r)
}