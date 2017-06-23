package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Sample Route",
		"GET",
		"/",
		MethodNotAllowed,
	},
	Route{
		"Sample Route",
		"GET",
		"/{[a-zA-Z0-9]}",
		MethodNotAllowed,
	},
        Route{
                "Health Check",
                "GET",
                "/v1/healthcheck",
                HealthCheck,
        },
}
