package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func middlewares(e *echo.Echo) *echo.Echo {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	// e.Use(middleware.Decompress())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	e.Use(cors())

	return e
}

func logger() echo.MiddlewareFunc {
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

func cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderAcceptEncoding,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAuthorization,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderOrigin,
			echo.HeaderXCSRFToken,
		},
		MaxAge: 2 * 3600,
	})
}
