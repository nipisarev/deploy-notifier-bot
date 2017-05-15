package main

import (
	"github.com/gin-gonic/gin"
	"./middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.LoggerResp())
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
