package tester_test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/stretchr/testify/assert"
)

func TestRequestTest(t *testing.T) {
	// setup
	e := echo.New()

	// testing
	code, body := tester.RequestTest(e, "METHOD", "/path", nil)

	// assertion
	assert.Equal(t, 404, code)
	assert.Equal(t, "{\"message\":\"Not Found\"}\n", body)
}

func TestRequestWithHeadersTest(t *testing.T) {
	// setup
	e := echo.New()

	// testing
	headers := map[string]string{"KEY": "VALUE"}
	code, body := tester.RequestWithHeadersTest(e, "METHOD", "/path", headers, nil)

	// assertion
	assert.Equal(t, 404, code)
	assert.Equal(t, "{\"message\":\"Not Found\"}\n", body)
}
