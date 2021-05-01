package app_test

import (
	"net/http"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErrorMethodNotAllowed(t *testing.T) {
	// setup
	testy := tester.New()

	// testing
	code, body := testy.RequestWithServe(http.MethodPost, "/", nil, nil)

	// assertion
	assert.Equal(t, http.StatusMethodNotAllowed, code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Method Not Allowed\",\"reason\":[]}\n", body)
}

func TestHTTPErrorNotFound(t *testing.T) {
	// setup
	testy := tester.New()

	// testing
	code, body := testy.RequestWithServe(http.MethodGet, "/xxx", nil, nil)

	// assertion
	assert.Equal(t, http.StatusNotFound, code)
	assert.Equal(t, "{\"error\":true,\"message\":\"Not Found\",\"reason\":[]}\n", body)
}
