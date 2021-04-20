package app

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/util/response"
)

func Injection(cfg *config.Config) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second
	e.HTTPErrorHandler = httpError

	e = middlewares(e, cfg)
	e = router(e)

	return e
}

func httpError(e error, c echo.Context) {
	code := http.StatusInternalServerError
	reason := make([]string, 0)

	if he, ok := e.(*echo.HTTPError); ok {
		if he.Code == http.StatusUnauthorized {
			msgErr := e.Error()
			reason = append(reason, getMessageError(msgErr))
		}
		code = he.Code
	}
	c.JSON(code, response.ErrorBody{
		Error:   true,
		Message: http.StatusText(code),
		Reason:  reason,
	})
}

func getMessageError(msg string) string {
	tmp := strings.Split(msg, ",")
	tmpMsg := ""
	if len(tmp) > 2 {
		tmpMsg = tmp[1]
		spl := strings.Split(tmpMsg, "=")
		tmpMsg = spl[1]
	}

	return tmpMsg
}
