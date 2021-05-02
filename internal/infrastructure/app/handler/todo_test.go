package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/shandysiswandi/echo-service/mocks"
	"github.com/shandysiswandi/echo-service/pkg/validation"
	"github.com/stretchr/testify/assert"
)

func TestFetchTodos(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", nil, nil)

		// mocks
		tdu := new(mocks.TodoUsecase)
		tdu.On("FetchTodos", context.TODO()).Return([]*domain.Todo{}, nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: tdu,
		}).FetchTodos(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", nil, nil)

		// mocks
		tdu := new(mocks.TodoUsecase)
		tdu.On("FetchTodos", context.TODO()).Return(nil, errors.New("error"))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: tdu,
		}).FetchTodos(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		tdu.AssertExpectations(t)
	})
}

func TestGetTodoById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")

		// mocks
		tdu := new(mocks.TodoUsecase)
		expected := domain.Todo{}
		actual := domain.Todo{}
		tdu.On("GetTodoByID", context.TODO(), "xxx-xxx-xxx").Return(&expected, nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: tdu,
		}).GetTodoById(c)

		// Assertions
		json.Unmarshal([]byte(rec.Body.Bytes()), &actual)
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected.ID, actual.ID)
		assert.Equal(t, expected.Title, actual.Title)
		assert.Equal(t, expected.Description, actual.Description)
		assert.Equal(t, expected.Completed, actual.Completed)
		assert.Equal(t, expected.Timestamp, actual.Timestamp)

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")

		// mocks
		tdu := new(mocks.TodoUsecase)
		tdu.On("GetTodoByID", context.TODO(), "xxx-xxx-xxx").Return(nil, errors.New("error"))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: tdu,
		}).GetTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		tdu.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoCreatePayload{
			Title:       "title of document todo",
			Description: "description of document todo",
		}
		tdu.On("CreateTodo", context.TODO(), payload).Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).CreateTodo(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error_Mock", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoCreatePayload{Title: "a", Description: "b"}
		tdu.On("CreateTodo", context.TODO(), payload).Return(errors.New("errLogic"))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).CreateTodo(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error_Bind", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b",}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: nil,
		}).CreateTodo(c)

		// Assertions
		assert.Equal(t, 206, len(h.Error()))
	})

	t.Run("Error_Validate", func(t *testing.T) {
		// setup
		paybody := `{"title":"","description":""}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		v := validation.New()

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: nil,
		}).CreateTodo(c)

		// Assertions
		assert.Equal(t, 219, len(h.Error()))
	})
}

func TestUpdateTodoById(t *testing.T) {
	t.Run("Success_One", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo", "completed":true}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoUpdatePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
			Completed:   true,
		}
		tdu.On("UpdateTodoByID", context.TODO(), payload).Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).UpdateTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Success_Two", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoUpdatePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
		}
		tdu.On("UpdateTodoByID", context.TODO(), payload).Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).UpdateTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Success_Three", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoUpdatePayload{
			ID:    "1",
			Title: "title of document todo",
		}
		tdu.On("UpdateTodoByID", context.TODO(), payload).Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).UpdateTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error_Mock", func(t *testing.T) {
		// setup
		paybody := `{"title":"12345","description":"12345 12345 12345"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoUpdatePayload{ID: "1", Title: "12345", Description: "12345 12345 12345"}
		tdu.On("UpdateTodoByID", context.TODO(), payload).Return(errors.New("errLogic"))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).UpdateTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error_Bind", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b",}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: nil,
		}).UpdateTodoById(c)

		// Assertions
		assert.Equal(t, 206, len(h.Error()))
	})

	t.Run("Error_Validate", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"a"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		v := validation.New()

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: nil,
		}).UpdateTodoById(c)

		// Assertions
		assert.Equal(t, 209, len(h.Error()))
	})
}

func TestReplaceTodoById(t *testing.T) {
	t.Run("Success_One", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo", "completed":true}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoReplacePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
			Completed:   true,
		}
		tdu.On("ReplaceTodoByID", context.TODO(), payload).Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).ReplaceTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Success_Two", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoReplacePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
		}
		tdu.On("ReplaceTodoByID", context.TODO(), payload).Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).ReplaceTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error_Mock", func(t *testing.T) {
		// setup
		paybody := `{"title":"12345","description":"12345 12345 12345"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		v := validation.New()

		// mocks
		tdu := new(mocks.TodoUsecase)
		payload := domain.TodoReplacePayload{ID: "1", Title: "12345", Description: "12345 12345 12345"}
		tdu.On("ReplaceTodoByID", context.TODO(), payload).Return(errors.New("errLogic"))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: tdu,
		}).ReplaceTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error_Bind", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b",}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: nil,
		}).ReplaceTodoById(c)

		// Assertions
		assert.Equal(t, 206, len(h.Error()))
	})

	t.Run("Error_Validate", func(t *testing.T) {
		// setup
		paybody := `{"title":"","description":""}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		v := validation.New()

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   v,
			TodoUsecase: nil,
		}).ReplaceTodoById(c)

		// Assertions
		assert.Equal(t, 221, len(h.Error()))
	})
}

func TestDeleteTodoById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		c, rec := tester.New().RequestWithContext(http.MethodDelete, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")

		// mocks
		tdu := new(mocks.TodoUsecase)
		tdu.On("DeleteTodoByID", context.TODO(), "xxx-xxx-xxx").Return(nil)

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: tdu,
		}).DeleteTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)

		// end mock
		tdu.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		c, _ := tester.New().RequestWithContext(http.MethodDelete, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")

		// mocks
		tdu := new(mocks.TodoUsecase)
		tdu.On("DeleteTodoByID", context.TODO(), "xxx-xxx-xxx").Return(errors.New("error"))

		// testing
		h := handler.New(handler.HandlerConfig{
			Validator:   nil,
			TodoUsecase: tdu,
		}).DeleteTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		tdu.AssertExpectations(t)
	})
}
