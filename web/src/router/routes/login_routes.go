package routes

import (
	"net/http"
	"web/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:    "/",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		RequireAuth: false,
	},
	{
		URI:    "/login",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		RequireAuth: false,
	},
}