package repository

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/port"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoMongo struct {
	db *mongo.Database
}

func NewTodoMongo(dbm *mongo.Database) port.TodoRepository {
	return &todoMongo{dbm}
}

func (tdm *todoMongo) Fetch(ctx context.Context) ([]*domain.Todo, error) {
	table := domain.Todo{}.TableName()
	results := make([]*domain.Todo, 0)
	filter := bson.D{}

	cur, err := tdm.db.Collection(table).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (tdm *todoMongo) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	result := &domain.Todo{}
	table := result.TableName()
	filter := bson.M{"id": id}

	err := tdm.db.Collection(table).FindOne(ctx, filter).Decode(result)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrTodoNotFound
	}

	return result, nil
}

func (tdm *todoMongo) Create(ctx context.Context, data domain.Todo) error {
	table := data.TableName()

	if _, err := tdm.db.Collection(table).InsertOne(ctx, &data); err != nil {
		return err
	}

	return nil
}

func (tdm *todoMongo) UpdateByID(ctx context.Context, data domain.Todo) error {
	table := data.TableName()
	filter := bson.M{"id": data.ID}
	update := bson.M{"$set": data}

	sr := tdm.db.Collection(table).FindOneAndUpdate(ctx, filter, update)
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		return sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		return domain.ErrTodoNotFound
	}

	return nil
}

func (tdm *todoMongo) ReplaceByID(ctx context.Context, data domain.Todo) error {
	table := data.TableName()
	filter := bson.M{"id": data.ID}

	sr := tdm.db.Collection(table).FindOneAndReplace(ctx, filter, data)
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		return sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		return domain.ErrTodoNotFound
	}

	return nil
}

func (tdm *todoMongo) DeleteByID(ctx context.Context, id string) error {
	table := domain.Todo{}.TableName()
	filter := bson.M{"id": id}

	sr := tdm.db.Collection(table).FindOneAndDelete(ctx, filter)
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		return sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		return domain.ErrTodoNotFound
	}

	return nil
}
