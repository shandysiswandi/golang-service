package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/middle"
)

func (h *Handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "welcome home")
}

func (h *Handler) Health(c echo.Context) error {
	return c.String(http.StatusOK, "health")
}

func (h *Handler) Graceful(c echo.Context) error {
	time.Sleep(1 * time.Second)
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) JWT(c echo.Context) error {
	accToken, _ := middle.CreateJWTToken("secret")
	return c.String(http.StatusOK, accToken)
}
