package app

import (
	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
)

func router(e *echo.Echo) *echo.Echo {
	// home
	homeHandler := handler.NewHomeHandler()
	e.GET("/", homeHandler.Home)
	e.GET("/graceful", homeHandler.Graceful)
	e.GET("/health", homeHandler.Health)

	return e
}
