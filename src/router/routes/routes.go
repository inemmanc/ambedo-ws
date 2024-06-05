package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Represents all API routes
//
// default route struct
type RouteDefault struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Insert all routes into the Router
func Configure(r *mux.Router) *mux.Router {
	var generalRoutes []RouteDefault
	generalRoutes = append(generalRoutes, usersRoutes...)

	for _, route := range generalRoutes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
