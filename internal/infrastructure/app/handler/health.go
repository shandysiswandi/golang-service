package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Health(c echo.Context) error {
	return c.String(http.StatusOK, "health")
}
