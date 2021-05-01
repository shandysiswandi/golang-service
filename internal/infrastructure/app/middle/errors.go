package middle

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/response"
)

func HTTPCustomError(e error, c echo.Context) {
	code := http.StatusInternalServerError
	reason := make([]string, 0)

	if he, ok := e.(*echo.HTTPError); ok {
		code = he.Code
		if he == middleware.ErrJWTMissing {
			reason = append(reason, "token not provide")
		}

		if he.Message == middleware.ErrJWTInvalid.Message {
			reason = append(reason, "token invalid or expired")
		}

		if he == echo.ErrStatusRequestEntityTooLarge {
			reason = append(reason, "request payload size over limit")
		}
	}

	c.JSON(code, response.ErrorBody{
		Error:   true,
		Message: http.StatusText(code),
		Reason:  reason,
	})
}
