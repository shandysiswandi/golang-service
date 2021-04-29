package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/adapter/repository"
	"github.com/shandysiswandi/echo-service/internal/adapter/usecase"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
)

func router(e *echo.Echo, validation *validator.Validate, dbm *mongodb.MongoDB) *echo.Echo {
	// register all repository
	tdr := repository.NewTodoMongo(dbm.GetDB())

	// register all usecase
	tdu := usecase.NewTodoUsecase(tdr)

	hc := handler.HandlerConfig{
		Validator:   validation,
		TodoUsecase: tdu,
	}
	h := handler.New(hc)

	e.GET("/", h.Home)
	e.GET("/graceful", h.Graceful)
	e.GET("/health", h.Health)

	e.GET("/todos", h.FetchTodos)
	e.GET("/todos/:id", h.GetTodoByID)
	e.POST("/todos", h.CreateTodo)
	e.PATCH("/todos/:id", h.UpdateTodoById)
	e.PUT("/todos/:id", h.ReplaceTodoById)
	e.DELETE("/todos/:id", h.DeleteTodoById)

	return e
}
