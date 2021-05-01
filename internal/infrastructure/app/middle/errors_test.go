package middle_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/middle"
	"github.com/stretchr/testify/assert"
)

func TestHTTPCustomError_ErrJWTMissing(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.Use(middle.JWT(""))

	// testing
	e.HTTPErrorHandler = middle.HTTPCustomError
	e.ServeHTTP(rec, req)

	// assertion
	assert.Equal(t, 400, rec.Code)
}

func TestHTTPCustomError_ErrJWTInvalid(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "Bearer xxx")
	rec := httptest.NewRecorder()
	e.Use(middle.JWT(""))

	// testing
	e.HTTPErrorHandler = middle.HTTPCustomError
	e.ServeHTTP(rec, req)

	// assertion
	assert.Equal(t, 401, rec.Code)
}

func TestHTTPCustomError_ErrStatusRequestEntityTooLarge(t *testing.T) {
	// setup
	hw := []byte("Hello, World! Lorem Ipsum Dolor Amet")
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Use(middleware.BodyLimit("2B"))

	// testing
	e.HTTPErrorHandler = middle.HTTPCustomError
	e.ServeHTTP(rec, req)

	// assertion
	assert.Equal(t, 413, rec.Code)
}
