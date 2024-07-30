package routes

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI string
	Method string 
	Function func(http.ResponseWriter, *http.Request)
	RequireAuth bool 
}

// SetRoutes sets all routes to the router
func SetRoutes(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(
				route.URI,
				middleware.Logger(middleware.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}