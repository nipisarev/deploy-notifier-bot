package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"context"
	"github.com/justinas/alice"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, wrap(route))
	}

	return router
}

func wrap(r Route) httprouter.Handle {
	return wrapHandler(
		alice.New(
			LogHandler,
		).
			Then(r.Handler))
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

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
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/:todoId",
		TodoShow,
	},
}
