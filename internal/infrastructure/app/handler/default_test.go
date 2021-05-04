package handler_test

import (
	"net/http"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	// setup
	testy := tester.New()
	c, rec := testy.RequestWithContext(http.MethodGet, "/", nil, nil)
	h, _ := testy.SetupHandlerTest()

	// testing
	assert.NoError(t, h.Home(c))

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "welcome home", rec.Body.String())
}

func TestHealth(t *testing.T) {
	// setup
	testy := tester.New()
	c, rec := testy.RequestWithContext(http.MethodGet, "/health", nil, nil)
	h, _ := testy.SetupHandlerTest()

	// testing
	assert.NoError(t, h.Health(c))

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "health", rec.Body.String())
}

func TestGraceful(t *testing.T) {
	// setup
	testy := tester.New()
	c, rec := testy.RequestWithContext(http.MethodGet, "/graceful", nil, nil)
	h, _ := testy.SetupHandlerTest()

	// testing
	assert.NoError(t, h.Graceful(c))

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}

func TestJWT(t *testing.T) {
	// setup
	testy := tester.New()
	c, rec := testy.RequestWithContext(http.MethodGet, "/jwt", nil, nil)
	h, _ := testy.SetupHandlerTest()

	// testing
	assert.NoError(t, h.JWT(c))

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}
