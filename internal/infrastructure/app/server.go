package app

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shandysiswandi/echo-service/internal/adapter/repository"
	"github.com/shandysiswandi/echo-service/internal/adapter/usecase"
	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/middle"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
	"github.com/shandysiswandi/echo-service/pkg/clock"
	"github.com/shandysiswandi/echo-service/pkg/gen"
	"github.com/shandysiswandi/echo-service/pkg/validation"
)

func New(cfg *config.Config, dbm *mongodb.MongoDB) *echo.Echo {
	e := echo.New()
	v := validation.New()

	/* ***** ***** ***** ***** ***** */
	/* setup server
	/* ***** ***** ***** ***** ***** */
	e.HideBanner = true
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second
	e.HTTPErrorHandler = middle.HTTPCustomError
	e.Validator = v

	/* ***** ***** ***** ***** ***** */
	/* setup middleware
	/* ***** ***** ***** ***** ***** */
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middle.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.Decompress())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	e.Use(middle.CORS())
	whiteList := []string{"/", "/xxx", "/health", "/jwt", "/graceful"}
	e.Use(middle.JWT(cfg.JWTSecret, whiteList...))

	/* ***** ***** ***** ***** ***** */
	/* setup router
	/* ***** ***** ***** ***** ***** */
	// register all library
	generator := gen.New()
	clk := clock.New()

	// register all repository
	tdr := repository.NewTodoMongo(dbm.GetDB())

	// register all usecase
	tdu := usecase.NewTodoUsecase(tdr)

	// register handler
	h := handler.New(handler.HandlerConfig{
		Validator:   v,
		Generator:   generator,
		Clock:       clk,
		TodoUsecase: tdu,
	})

	e.GET("/", h.Home)
	e.GET("/graceful", h.Graceful)
	e.GET("/health", h.Health)
	e.GET("/jwt", h.JWT)

	e.GET("/todos", h.FetchTodos)
	e.GET("/todos/:id", h.GetTodoById)
	e.POST("/todos", h.CreateTodo)
	e.PATCH("/todos/:id", h.UpdateTodoById)
	e.PUT("/todos/:id", h.ReplaceTodoById)
	e.DELETE("/todos/:id", h.DeleteTodoById)

	return e
}
