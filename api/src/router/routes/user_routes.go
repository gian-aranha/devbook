package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.GetUsers,
		RequireAuth: true,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.GetUser,
		RequireAuth: true,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI: "/users/{userId}/follow",
		Method: http.MethodPost,
		Function: controllers.FollowUser,
		RequireAuth: true,
	},
}