package main

import (
	"net/http"
	"./models"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "welcome")
}

func TodoIndex(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=UTF-8")

	c.JSON(http.StatusOK, todos)
}

func TodoShow(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=UTF-8") //.Header().Set("Content-Type", "application/json; charset=UTF-8")

	i, err := strconv.Atoi(c.Param("todoId"))
	if err != nil {
		panic(err)
	}
	todo, err := RepoFindTodod(i);
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todo)
}

func TodoCreate(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	t := RepoCreateTodo(todo)
	c.Header("Content-Type", "application/json; charset=UTF-8")
	c.JSON(http.StatusCreated, t)
}
