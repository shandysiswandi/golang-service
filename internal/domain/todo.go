package domain

import (
	"errors"
	"time"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type Todo struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Completed   bool      `json:"completed" bson:"completed,omitempty"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp,omitempty"`
}

type TodoCreatePayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TodoUpdatePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"omitempty"`
	Description string `json:"description" validate:"omitempty"`
	Completed   bool   `json:"completed" validate:"omitempty"`
}

type TodoReplacePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed" validate:"required"`
}

func (Todo) TableName() string {
	return "todos"
}
