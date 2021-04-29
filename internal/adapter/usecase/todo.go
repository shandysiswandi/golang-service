package usecase

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/port"
)

type todoUsecase struct {
	todoRepo port.TodoRepository
}

func NewTodoUsecase(tdu port.TodoRepository) port.TodoUsecase {
	return &todoUsecase{tdu}
}

func (tdu *todoUsecase) FetchTodos(context.Context) ([]*domain.Todo, error) {
	return nil, nil
}

func (tdu *todoUsecase) GetTodoByID(context.Context, string) (*domain.Todo, error) {
	return nil, nil
}

func (tdu *todoUsecase) CreateTodo(context.Context, domain.TodoCreatePayload) error {
	return nil
}

func (tdu *todoUsecase) UpdateTodoByID(context.Context, domain.TodoUpdatePayload) error {
	return nil
}

func (tdu *todoUsecase) ReplaceTodoByID(context.Context, domain.TodoReplacePayload) error {
	return nil
}

func (tdu *todoUsecase) DeleteTodoByID(context.Context, string) error {
	return nil
}
