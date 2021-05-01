package middle

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	var (
		times   = "\033[33mTIME\033[0m: ${time_rfc3339}"
		method  = "\033[33mMETHOD\033[0m: ${method}"
		uri     = "\033[33mURI\033[0m: ${uri}"
		ip      = "\033[33mIP\033[0m: ${remote_ip}"
		status  = "\033[33mSTATUS\033[0m: ${status}"
		err     = "\033[33mERROR\033[0m: ${error}"
		latency = "\033[33mLATENCY\033[0m: ${latency_human}"
	)

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s\n", times, method, uri, ip, status, err, latency),
	})
}
