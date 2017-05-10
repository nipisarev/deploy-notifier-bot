package main

import (
	"net/http"
	"fmt"
	"html"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

func FetchParams(req *http.Request) httprouter.Params {
	ctx := req.Context()
	return ctx.Value("params").(httprouter.Params)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %q", html.EscapeString(r.RequestURI))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation", Completed: true},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	params := FetchParams(r)
	fmt.Fprintf(w, "Todo is %s", params.ByName("todoId"))
}
