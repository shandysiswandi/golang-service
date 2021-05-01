package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "welcome home")
}

func (h *handler) Health(c echo.Context) error {
	return c.String(http.StatusOK, "health")
}

func (h *handler) Graceful(c echo.Context) error {
	time.Sleep(1 * time.Second)
	return c.String(http.StatusOK, "OK")
}
