package port

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/domain"
)

type TodoUsecase interface {
	FetchTodos(context.Context) ([]*domain.Todo, error)
	GetTodoByID(context.Context, string) (*domain.Todo, error)

	CreateTodo(context.Context, domain.TodoCreatePayload) error
	UpdateTodoByID(context.Context, domain.TodoUpdatePayload) error
	ReplaceTodoByID(context.Context, domain.TodoReplacePayload) error
	DeleteTodoByID(context.Context, string) error
}

type TodoRepository interface {
	Fetch(context.Context) ([]*domain.Todo, error)
	GetByID(context.Context, string) (*domain.Todo, error)

	Create(context.Context, domain.Todo) error
	UpdateByID(context.Context, domain.Todo) error
	ReplaceByID(context.Context, domain.Todo) error
	DeleteByID(context.Context, string) error
}
