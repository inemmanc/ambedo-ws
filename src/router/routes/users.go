package routes

import "net/http"

var usersRoutes = []RouteDefault{
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/users{userID}",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     func(w http.ResponseWriter, r http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/users{userID}",
		Method:       http.MethodPut,
		Function:     func(w http.ResponseWriter, r http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/users{userID}",
		Method:       http.MethodDelete,
		Function:     func(w http.ResponseWriter, r http.Request) {},
		AuthRequired: false,
	},
}
