package app_test

import (
	"net/http"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErrorMethodNotAllowed(t *testing.T) {
	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)

	// testing
	code, body := tester.RequestTest(e, http.MethodPost, "/", nil)

	// assertion
	assert.Equal(t, http.StatusMethodNotAllowed, code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Method Not Allowed\",\"reason\":[]}\n", body)
}

func TestHTTPErrorNotFound(t *testing.T) {
	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)

	// testing
	code, body := tester.RequestTest(e, http.MethodGet, "/xxx", nil)

	// assertion
	assert.Equal(t, http.StatusNotFound, code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Not Found\",\"reason\":[]}\n", body)
}

func TestHTTPErrorUnauthorized(t *testing.T) {
	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)

	// testing
	headers := map[string]string{"Authorization": "Bearer fake-jwt"}
	code, body := tester.RequestWithHeadersTest(e, http.MethodGet, "/jwt", headers, nil)

	// assertion
	assert.Equal(t, http.StatusUnauthorized, code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Unauthorized\",\"reason\":[\"token invalid or expired\"]}\n", body)
}

func TestHTTPErrorBadRequest(t *testing.T) {
	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)

	// testing
	headers := map[string]string{"Authorization": "Bearer "}
	code, body := tester.RequestWithHeadersTest(e, http.MethodGet, "/jwt", headers, nil)

	// assertion
	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Bad Request\",\"reason\":[\"token not provide\"]}\n", body)
}

// func TestHTTPErrorInternalServerError(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
// 	rec := httptest.NewRecorder()

// 	token := "a"
// 	req.Header.Set("Authorization", "Bearer "+token)

// 	cfg := &config.Config{JWTSecret: ""}
// 	e := app.Injection(cfg)

// 	e.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	assert.Equal(t, "{\"error\":true,\"message\":\"Bad Request\",\"reason\":[\"missing or malformed jwt\"]}\n", rec.Body.String())
// }
