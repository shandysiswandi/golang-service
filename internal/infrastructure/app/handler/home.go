package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/util/response"
)

type homeHandler struct{}

func NewHomeHandler() *homeHandler {
	return &homeHandler{}
}

func (hh *homeHandler) Home(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, response.Success("welcome home", []string{}))
}

func (hh *homeHandler) Graceful(ctx echo.Context) error {
	time.Sleep(2 * time.Second)
	return ctx.JSON(http.StatusOK, "OK")
}

func (hh *homeHandler) Health(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "health")
}
