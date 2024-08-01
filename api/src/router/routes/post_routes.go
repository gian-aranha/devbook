package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:    "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI:    "/posts",
		Method: http.MethodGet,
		Function: controllers.GetPosts,
		RequireAuth: true,
	},
	{
		URI:    "/posts/{postId}",
		Method: http.MethodGet,
		Function: controllers.GetPost,
		RequireAuth: true,
	},
	{
		URI:    "/posts/{postId}",
		Method: http.MethodPut,
		Function: controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI:    "/posts/{postId}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		RequireAuth: true,
	},
	{
		URI: "/users/{userId}/posts",
		Method: http.MethodGet,
		Function: controllers.GetUserPosts,
		RequireAuth: true,
	},
	{
		URI: "/posts/{postId}/like",
		Method: http.MethodPost,
		Function: controllers.LikePost,
		RequireAuth: true,
	},
	{
		URI: "/posts/{postId}/unlike",
		Method: http.MethodPost,
		Function: controllers.UnlikePost,
		RequireAuth: true,
	},
}