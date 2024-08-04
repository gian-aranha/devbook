package router

import (
	"web/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with all routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.SetRoutes(r)
}