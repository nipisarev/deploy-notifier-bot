package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"html"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Welcome, %q", html.EscapeString(r.RequestURI))
}

func TodoIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write presentation", Completed: true},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo, %s", ps.ByName("todoId"))
}
