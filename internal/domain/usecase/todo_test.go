package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/domain/port"
	"github.com/shandysiswandi/echo-service/internal/domain/usecase"
	"github.com/shandysiswandi/echo-service/mocks"
	"github.com/stretchr/testify/assert"
)

func setup() (context.Context, *mocks.TodoRepository, port.TodoUsecase) {
	ctx := context.Background()
	tRepo := new(mocks.TodoRepository)
	// idgen := new(mocks.IDGenerator)
	// ts := new(mocks.Times)
	use := usecase.NewTodoUsecase(tRepo)

	return ctx, tRepo, use
}

func TestNewTodoUsecase(t *testing.T) {
	assert.NotNil(t, usecase.NewTodoUsecase(nil))
}

func TestFetchTodos(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		actual := []*domain.Todo{}
		repo.On("Fetch", ctx).Return(actual, nil)

		// testing
		data, err := use.FetchTodos(ctx)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, actual, data)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		actual := []*domain.Todo{}
		repo.On("Fetch", ctx).Return(actual, errors.New("err"))

		// testing
		_, err := use.FetchTodos(ctx)

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestGetTodoByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		actual := &domain.Todo{}
		repo.On("GetByID", ctx, "xxx").Return(actual, nil)

		// testing
		data, err := use.GetTodoByID(ctx, "xxx")

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, actual, data)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		repo.On("GetByID", ctx, "xxx").Return(nil, errors.New("err"))

		// testing
		_, err := use.GetTodoByID(ctx, "xxx")

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()

		payload := domain.TodoCreatePayload{ID: "1", Title: "a", Description: "b", Timestamp: time.Now()}
		data := domain.Todo{
			ID:          payload.ID,
			Title:       payload.Title,
			Description: payload.Description,
			Completed:   false,
			Timestamp:   payload.Timestamp,
		}
		repo.On("Create", ctx, data).Return(nil)

		// testing
		err := use.CreateTodo(ctx, payload)

		// assertion
		assert.NoError(t, err)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()

		pld := domain.TodoCreatePayload{}
		data := domain.Todo{ID: pld.ID, Title: pld.Title, Description: pld.Description, Completed: false, Timestamp: pld.Timestamp}
		repo.On("Create", ctx, data).Return(errors.New("err"))

		// testing
		err := use.CreateTodo(ctx, pld)

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestUpdateTodoByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		tup := domain.TodoUpdatePayload{}
		td := domain.Todo{}
		repo.On("UpdateByID", ctx, td).Return(nil)

		// testing
		err := use.UpdateTodoByID(ctx, tup)

		// assertion
		assert.NoError(t, err)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		tup := domain.TodoUpdatePayload{}
		td := domain.Todo{}
		repo.On("UpdateByID", ctx, td).Return(errors.New("err"))

		// testing
		err := use.UpdateTodoByID(ctx, tup)

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestReplaceTodoByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		trp := domain.TodoReplacePayload{}
		td := domain.Todo{}
		repo.On("ReplaceByID", ctx, td).Return(nil)

		// testing
		err := use.ReplaceTodoByID(ctx, trp)

		// assertion
		assert.NoError(t, err)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		trp := domain.TodoReplacePayload{}
		td := domain.Todo{}
		repo.On("ReplaceByID", ctx, td).Return(errors.New("err"))

		// testing
		err := use.ReplaceTodoByID(ctx, trp)

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestDeleteTodoByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		repo.On("DeleteByID", ctx, "xxx").Return(nil)

		// testing
		err := use.DeleteTodoByID(ctx, "xxx")

		// assertion
		assert.NoError(t, err)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		ctx, repo, use := setup()
		repo.On("DeleteByID", ctx, "xxx").Return(errors.New("err"))

		// testing
		err := use.DeleteTodoByID(ctx, "xxx")

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}
