package main

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/hosts",
		HostIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/hosts",
		AddHost,
	},
	Route{
		"GetHost",
		"GET",
		"/hosts/:hostName",
		GetHost,
	},
}
