package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	cfg := &config.Config{JWTSecret: ""}
	e := app.New(cfg)

	c := e.NewContext(req, rec)

	hc := handler.HandlerConfig{}
	h := handler.New(hc)

	expBody := "health"

	// Assertions
	assert.NoError(t, h.Health(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expBody, rec.Body.String())
}
