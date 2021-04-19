package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/stretchr/testify/assert"
)

func Test_HTTPError_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	e := app.Injection(nil)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Method Not Allowed\",\"reason\":[]}\n", rec.Body.String())
}

func Test_HTTPError_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/xxx", nil)
	rec := httptest.NewRecorder()

	e := app.Injection(nil)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Not Found\",\"reason\":[]}\n", rec.Body.String())
}
