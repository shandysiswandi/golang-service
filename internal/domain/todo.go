package domain

import (
	"errors"
	"time"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type (
	Todo struct {
		ID          string    `json:"id" bson:"id"`
		Title       string    `json:"title" bson:"title,omitempty"`
		Description string    `json:"description" bson:"description,omitempty"`
		Completed   bool      `json:"completed" bson:"completed,omitempty"`
		Timestamp   time.Time `json:"timestamp" bson:"timestamp,omitempty"`
	}

	TodoCreatePayload struct {
		ID          string    `json:"id"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Timestamp   time.Time `json:"timestamp"`
	}

	TodoUpdatePayload struct {
		ID          string `json:"id"`
		Title       string `json:"title" validate:"omitempty,min=5"`
		Description string `json:"description" validate:"omitempty,min=15"`
		Completed   bool   `json:"completed" validate:"omitempty"`
	}

	TodoReplacePayload struct {
		ID          string `json:"id"`
		Title       string `json:"title" validate:"required,min=5"`
		Description string `json:"description" validate:"required,min=15"`
		Completed   bool   `json:"completed" validate:"omitempty"`
	}
)

func (Todo) TableName() string {
	return "todos"
}
