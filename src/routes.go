package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"context"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		//handler := route.Handle
		//handler = Logger(handler, route.Name)

		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	return router
}

//func wrap(p string, h func(http.ResponseWriter, *http.Request)) (string, httprouter.Handle) {
//	return p, wrapHandler(alice.New().ThenFunc(h))
//}
//
//func wrapHandler(h http.Handler) httprouter.Handle {
//	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//		ctx := context.WithValue(r.Context(), "params", ps)
//		r = r.WithContext(ctx)
//		h.ServeHTTP(w, r)
//	}
//}


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
