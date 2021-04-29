package repository

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/port"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoMongo struct {
	dbm *mongo.Database
}

func NewTodoMongo(dbm *mongo.Database) port.TodoRepository {
	return &todoMongo{dbm}
}

func (tdm *todoMongo) Fetch(context.Context) ([]*domain.Todo, error) {
	return nil, nil
}

func (tdm *todoMongo) GetByID(context.Context, string) (*domain.Todo, error) {
	return nil, nil
}

func (tdm *todoMongo) Create(context.Context, domain.Todo) error {
	return nil
}

func (tdm *todoMongo) UpdateByID(context.Context, domain.Todo) error {
	return nil
}

func (tdm *todoMongo) ReplaceByID(context.Context, domain.Todo) error {
	return nil
}

func (tdm *todoMongo) DeleteByID(context.Context, string) error {
	return nil
}
