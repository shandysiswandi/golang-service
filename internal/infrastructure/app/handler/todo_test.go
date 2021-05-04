package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/domain"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/stretchr/testify/assert"
)

func TestFetchTodos(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", nil, nil)
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		ret.TodoUsecase.On("FetchTodos", ctx).Return([]*domain.Todo{}, nil)
		h := hen.FetchTodos(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", nil, nil)
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		ret.TodoUsecase.On("FetchTodos", ctx).Return(nil, errors.New("error"))
		h := hen.FetchTodos(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})
}

func TestGetTodoById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		expected := domain.Todo{}
		actual := domain.Todo{}
		ret.TodoUsecase.On("GetTodoByID", ctx, "xxx-xxx-xxx").Return(&expected, nil)

		// testing
		h := hen.GetTodoById(c)

		// Assertions
		json.Unmarshal([]byte(rec.Body.Bytes()), &actual)
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected.ID, actual.ID)

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		ret.TodoUsecase.On("GetTodoByID", ctx, "xxx-xxx-xxx").Return(nil, errors.New("error"))

		// testing
		h := hen.GetTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {
	tm := time.Now()
	id := "id123"
	testy := tester.New()
	hen, ret := testy.SetupHandlerTest()

	t.Run("Success", func(t *testing.T) {
		// setup
		paybody := `{"title":"title","description":"description of document"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoCreatePayload{
			ID:          id,
			Title:       "title",
			Description: "description of document",
			Timestamp:   tm,
		}
		ret.TodoUsecase.On("CreateTodo", ctx, payload).Return(nil)
		ret.Generator.On("Nanoid").Return(id)
		ret.Clocker.On("Now").Return(tm)

		// testing
		h := hen.CreateTodo(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
		ret.Generator.AssertExpectations(t)
		ret.Clocker.AssertExpectations(t)
	})

	t.Run("Error_Mock", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoCreatePayload{ID: id, Title: "a", Description: "b", Timestamp: tm}
		ret.TodoUsecase.On("CreateTodo", ctx, payload).Return(errors.New("errLogic"))
		ret.Generator.On("Nanoid").Return(id)
		ret.Clocker.On("Now").Return(tm)

		// testing
		h := hen.CreateTodo(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
		ret.Generator.AssertExpectations(t)
		ret.Clocker.AssertExpectations(t)
	})

	t.Run("Error_Bind", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b",}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))

		// testing
		h := hen.CreateTodo(c)

		// Assertions
		assert.Equal(t, 206, len(h.Error()))
	})

	t.Run("Error_Validate", func(t *testing.T) {
		// setup
		paybody := `{"title":"","description":""}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))

		// testing
		h := hen.CreateTodo(c)

		// Assertions
		assert.Equal(t, 219, len(h.Error()))
	})
}

func TestUpdateTodoById(t *testing.T) {
	t.Run("Success_One", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo", "completed":true}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoUpdatePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
			Completed:   true,
		}
		ret.TodoUsecase.On("UpdateTodoByID", ctx, payload).Return(nil)

		// testing
		h := hen.UpdateTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Success_Two", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoUpdatePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
		}
		ret.TodoUsecase.On("UpdateTodoByID", ctx, payload).Return(nil)

		// testing
		h := hen.UpdateTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Success_Three", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoUpdatePayload{
			ID:    "1",
			Title: "title of document todo",
		}
		ret.TodoUsecase.On("UpdateTodoByID", ctx, payload).Return(nil)

		// testing
		h := hen.UpdateTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error_Mock", func(t *testing.T) {
		// setup
		paybody := `{"title":"12345","description":"12345 12345 12345"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoUpdatePayload{ID: "1", Title: "12345", Description: "12345 12345 12345"}
		ret.TodoUsecase.On("UpdateTodoByID", ctx, payload).Return(errors.New("errLogic"))

		// testing
		h := hen.UpdateTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error_Bind", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b",}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		hen, _ := testy.SetupHandlerTest()

		// testing
		h := hen.UpdateTodoById(c)

		// Assertions
		assert.Equal(t, 206, len(h.Error()))
	})

	t.Run("Error_Validate", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"a"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		hen, _ := testy.SetupHandlerTest()

		// testing
		h := hen.UpdateTodoById(c)

		// Assertions
		assert.Equal(t, 209, len(h.Error()))
	})
}

func TestReplaceTodoById(t *testing.T) {
	t.Run("Success_One", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo", "completed":true}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, rec := tester.New().RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoReplacePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
			Completed:   true,
		}
		ret.TodoUsecase.On("ReplaceTodoByID", ctx, payload).Return(nil)

		// testing
		h := hen.ReplaceTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Success_Two", func(t *testing.T) {
		// setup
		paybody := `{"title":"title of document todo","description":"description of document todo"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoReplacePayload{
			ID:          "1",
			Title:       "title of document todo",
			Description: "description of document todo",
		}
		ret.TodoUsecase.On("ReplaceTodoByID", ctx, payload).Return(nil)

		// testing
		h := hen.ReplaceTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK\"\n", rec.Body.String())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error_Mock", func(t *testing.T) {
		// setup
		paybody := `{"title":"12345","description":"12345 12345 12345"}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		payload := domain.TodoReplacePayload{ID: "1", Title: "12345", Description: "12345 12345 12345"}
		ret.TodoUsecase.On("ReplaceTodoByID", ctx, payload).Return(errors.New("errLogic"))

		// testing
		h := hen.ReplaceTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error_Bind", func(t *testing.T) {
		// setup
		paybody := `{"title":"a","description":"b",}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		hen, _ := testy.SetupHandlerTest()

		// testing
		h := hen.ReplaceTodoById(c)

		// Assertions
		assert.Equal(t, 206, len(h.Error()))
	})

	t.Run("Error_Validate", func(t *testing.T) {
		// setup
		paybody := `{"title":"","description":""}`
		header := map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodGet, "/todos", header, strings.NewReader(paybody))
		hen, _ := testy.SetupHandlerTest()

		// testing
		h := hen.ReplaceTodoById(c)

		// Assertions
		assert.Equal(t, 221, len(h.Error()))
	})
}

func TestDeleteTodoById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		testy := tester.New()
		c, rec := testy.RequestWithContext(http.MethodDelete, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		ret.TodoUsecase.On("DeleteTodoByID", ctx, "xxx-xxx-xxx").Return(nil)

		// testing
		h := hen.DeleteTodoById(c)

		// Assertions
		assert.NoError(t, h)
		assert.Equal(t, http.StatusOK, rec.Code)

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		testy := tester.New()
		c, _ := testy.RequestWithContext(http.MethodDelete, "/todos", nil, nil)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("xxx-xxx-xxx")
		hen, ret := testy.SetupHandlerTest()
		ctx := c.Request().Context()

		// mocks
		ret.TodoUsecase.On("DeleteTodoByID", ctx, "xxx-xxx-xxx").Return(errors.New("error"))

		// testing
		h := hen.DeleteTodoById(c)

		// Assertions
		assert.Equal(t, "code=500, message=Internal Server Error", h.Error())

		// end mock
		ret.TodoUsecase.AssertExpectations(t)
	})
}
