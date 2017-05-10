package main

import (
	"net/http"
	"time"
	"log"
)

type logHandler struct {
	handler http.Handler
}

func (h logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	h.handler.ServeHTTP(w, r)
	params := FetchParams(r)

	log.Printf("%fs\t%s\t%s\tparams: %s", time.Since(start).Seconds(), r.Method, r.RequestURI, params)
}

func LogHandler(h http.Handler) http.Handler {
	return logHandler{h}
}
