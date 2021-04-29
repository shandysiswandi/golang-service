package domain

import "time"

type Todo struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Completed   bool      `json:"completed" bson:"completed"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
}

type TodoCreatePayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TodoUpdatePayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed" validate:"required"`
}

type TodoReplacePayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed" validate:"required"`
}

func (Todo) TableName() string {
	return "todos"
}
