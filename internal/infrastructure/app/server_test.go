package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/stretchr/testify/assert"
)

func Test_HTTPError_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	cfg := &config.Config{JWTSecret: ""}
	e := app.New(cfg)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Method Not Allowed\",\"reason\":[]}\n", rec.Body.String())
}

func Test_HTTPError_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/xxx", nil)
	rec := httptest.NewRecorder()

	cfg := &config.Config{JWTSecret: ""}
	e := app.New(cfg)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Not Found\",\"reason\":[]}\n", rec.Body.String())
}

func Test_HTTPError_Unauthorized(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/jwt", nil)
	rec := httptest.NewRecorder()

	token := "fake-jwt"
	req.Header.Set("Authorization", "Bearer "+token)

	cfg := &config.Config{JWTSecret: ""}
	e := app.New(cfg)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Unauthorized\",\"reason\":[\"token invalid or expired\"]}\n", rec.Body.String())
}

func Test_HTTPError_BadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/jwt", nil)
	rec := httptest.NewRecorder()

	token := ""
	req.Header.Set("Authorization", "Bearer "+token)

	cfg := &config.Config{JWTSecret: ""}
	e := app.New(cfg)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Bad Request\",\"reason\":[\"token not provide\"]}\n", rec.Body.String())
}

// func Test_HTTPError_InternalServerError(t *testing.T) {
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
