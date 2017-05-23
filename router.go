package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nipisarev/deploy-notifier-bot/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.Monger())
	router.Use(middlewares.Logger())
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
