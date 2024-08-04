package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all web application routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// SetRoutes sets all routes to the router
func SetRoutes(r *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}