package router

import (
	"ambedo-ws/src/router/routes"

	"github.com/gorilla/mux"
)

// Returns router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
