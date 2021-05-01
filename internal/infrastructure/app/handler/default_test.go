package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)
	c := e.NewContext(req, rec)

	// testing
	hc := handler.HandlerConfig{}
	h := handler.New(hc)

	// Assertions
	assert.NoError(t, h.Home(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "welcome home", rec.Body.String())
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)
	c := e.NewContext(req, rec)

	// testing
	hc := handler.HandlerConfig{}
	h := handler.New(hc)

	// Assertions
	assert.NoError(t, h.Health(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "health", rec.Body.String())
}

func TestGraceful(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/graceful", nil)
	rec := httptest.NewRecorder()

	// setup
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)
	c := e.NewContext(req, rec)

	// testing
	hc := handler.HandlerConfig{}
	h := handler.New(hc)

	// Assertions
	assert.NoError(t, h.Graceful(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}
