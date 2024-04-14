package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.GetAllUser,
		RequiredAuth: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.GetUserById,
		RequiredAuth: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequiredAuth: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequiredAuth: false,
	},
}