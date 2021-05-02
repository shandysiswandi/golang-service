package usecase

import (
	"context"
	"time"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/port"
	"github.com/shandysiswandi/echo-service/pkg/nanoid"
)

type todoUsecase struct {
	todoRepo    port.TodoRepository
	idGenerator nanoid.IDGenerator
}

func NewTodoUsecase(tdu port.TodoRepository, idg nanoid.IDGenerator) port.TodoUsecase {
	return &todoUsecase{tdu, idg}
}

func (tdu *todoUsecase) FetchTodos(ctx context.Context) ([]*domain.Todo, error) {
	return tdu.todoRepo.Fetch(ctx)
}

func (tdu *todoUsecase) GetTodoByID(ctx context.Context, id string) (*domain.Todo, error) {
	return tdu.todoRepo.GetByID(ctx, id)
}

func (tdu *todoUsecase) CreateTodo(ctx context.Context, payload domain.TodoCreatePayload) error {
	data := domain.Todo{
		ID:          tdu.idGenerator.Generate(),
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   false,
		Timestamp:   time.Now(),
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
