package app

import (
	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
)

func router(e *echo.Echo) *echo.Echo {

	hc := handler.HandlerConfig{}
	h := handler.New(hc)

	e.GET("/", h.Home)
	e.GET("/graceful", h.Graceful)
	e.GET("/health", h.Health)

	return e
}
