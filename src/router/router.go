package router

import "github.com/gorilla/mux"

// Returns router with configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
