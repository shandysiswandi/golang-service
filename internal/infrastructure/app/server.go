package app

import (
	"fmt"
	"net/http"
	"time"

	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/response"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
	"github.com/shandysiswandi/echo-service/internal/util/is"
)

type (
	UserJWT struct {
		ID            int    `json:"id"`
		Email         string `json:"email"`
		CompanyID     int    `json:"company_id"`
		UserCompanyID int    `json:"user_company_id"`
	}

	JWTClaim struct {
		jwtlib.StandardClaims
		SessID         int         `json:"sess_id"`
		User           UserJWT     `json:"user"`
		Lang           interface{} `json:"lang"`
		SessionSetting int         `json:"session_setting"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func New(cfg *config.Config, dbm *mongodb.MongoDB) *echo.Echo {
	e := echo.New()
	validation := validator.New()

	// setup server
	e.HideBanner = true
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second
	e.HTTPErrorHandler = httpError
	e.Validator = &CustomValidator{validator: validation}

	// setup middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))
	// e.Use(middleware.Decompress())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	e.Use(cors())
	// e.Use(jwt(cfg.JWTSecret))

	// setup router
	e = router(e, validation, dbm)

	return e
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}

func httpError(e error, c echo.Context) {
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

func jwt(key string) echo.MiddlewareFunc {
	jwtCfg := middleware.JWTConfig{
		Claims:     &JWTClaim{},
		SigningKey: []byte(key),
		Skipper: func(c echo.Context) bool {
			whiteList := []string{"/", "/xxx", "/graceful", "/health", "/gzip"}
			path := c.Request().URL.Path
			return is.InArrayString(whiteList, path)
		},
	}
	return middleware.JWTWithConfig(jwtCfg)
}
