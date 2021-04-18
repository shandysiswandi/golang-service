package app

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/config"
)

func Injection(cfg *config.Config) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	e = middlewares(e)
	e = router(e)

	return e
}
