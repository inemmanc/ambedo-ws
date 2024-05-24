package routes

import "net/http"

// Represents all API routes
//
// default route struct
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, http.Request)
	AuthRequired bool
}
