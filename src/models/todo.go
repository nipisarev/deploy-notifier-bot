package models

import "time"

type Todo struct {
	Id        int	    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
