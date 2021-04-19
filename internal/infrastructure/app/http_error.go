package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/util/response"
)

func httpError(e error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := e.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.JSON(code, response.ErrorBody{
		Error:   true,
		Message: http.StatusText(code),
		Reason:  []int{},
	})
}
