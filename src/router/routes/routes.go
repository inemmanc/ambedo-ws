package routes

import "net/http"

// Represents all API routes
//
// default route struct
type RouteDefault struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, http.Request)
	AuthRequired bool
}
