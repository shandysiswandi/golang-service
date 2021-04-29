package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/response"
)

func (h *handler) Home(c echo.Context) error {
	return c.JSON(http.StatusOK, response.Success("welcome home", []string{}))
}
