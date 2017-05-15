package main

import (
	"errors"
	"fmt"
	"./models"
)

var currentId int
var todos models.Todos

func init() {
	RepoCreateTodo(models.Todo{Name: "Write presentation"})
	RepoCreateTodo(models.Todo{Name: "Host meetup"})
}

func RepoCreateTodo(t models.Todo) models.Todo {
	currentId++
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoFindTodod(id int) (models.Todo, error) {
	for _, t := range todos {
		if t.Id == id {
			return t, nil
		}
	}

	return models.Todo{}, errors.New("not found Todo")
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
