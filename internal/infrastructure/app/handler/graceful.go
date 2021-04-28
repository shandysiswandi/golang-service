package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handler) Graceful(c echo.Context) error {
	time.Sleep(2 * time.Second)
	return c.JSON(http.StatusOK, "OK")
}
