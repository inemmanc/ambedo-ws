package routes

import "net/http"

var usersRoutes = []RouteDefault{
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r http.Request) {},
		AuthRequired: false,
	},
}
