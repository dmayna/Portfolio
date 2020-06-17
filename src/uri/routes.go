package uri

import (
	"net/http"

	"uri/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HandleHome",
		"GET",
		"/",
		handlers.HandleHome,
	},
	Route{
		"HandleProjects",
		"GET",
		"/projects",
		handlers.HandleProjects,
	},
	Route{
		"HandleResume",
		"GET",
		"/resume",
		handlers.HandleResume,
	},
	Route{
		"HandleSports",
		"GET",
		"/sports",
		handlers.HandleSports,
	},
	Route{
		"HandleMusic",
		"GET",
		"/music",
		handlers.HandleMusic,
	},
	Route{
		"HandleContact",
		"GET",
		"/contact",
		handlers.HandleContact,
	},
	Route{
		"HandleBlockchain",
		"GET",
		"/blockchain",
		handlers.HandleBlockchain,
	},
	Route{
		"HandleStart",
		"GET",
		"/start",
		handlers.HandleStart,
	},
}
