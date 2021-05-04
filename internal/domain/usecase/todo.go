package usecase

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/domain/port"
)

type todoUsecase struct {
	todoRepo port.TodoRepository
}

func NewTodoUsecase(tdu port.TodoRepository) port.TodoUsecase {
	return &todoUsecase{tdu}
}

func (tdu *todoUsecase) FetchTodos(ctx context.Context) ([]*domain.Todo, error) {
	return tdu.todoRepo.Fetch(ctx)
}

func (tdu *todoUsecase) GetTodoByID(ctx context.Context, id string) (*domain.Todo, error) {
	return tdu.todoRepo.GetByID(ctx, id)
}

func (tdu *todoUsecase) CreateTodo(ctx context.Context, payload domain.TodoCreatePayload) error {
	data := domain.Todo{
		ID:          payload.ID,
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   false,
		Timestamp:   payload.Timestamp,
	}
	return tdu.todoRepo.Create(ctx, data)
}

func (tdu *todoUsecase) UpdateTodoByID(ctx context.Context, payload domain.TodoUpdatePayload) error {
	data := domain.Todo{
		ID:          payload.ID,
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   payload.Completed,
	}
	return tdu.todoRepo.UpdateByID(ctx, data)
}

func (tdu *todoUsecase) ReplaceTodoByID(ctx context.Context, payload domain.TodoReplacePayload) error {
	data := domain.Todo{
		ID:          payload.ID,
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   payload.Completed,
	}
	return tdu.todoRepo.ReplaceByID(ctx, data)
}

func (tdu *todoUsecase) DeleteTodoByID(ctx context.Context, id string) error {
	return tdu.todoRepo.DeleteByID(ctx, id)
}
