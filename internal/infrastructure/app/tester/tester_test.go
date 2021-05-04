package tester_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/tester"
	"github.com/stretchr/testify/assert"
)

func TestRequestWithServe(t *testing.T) {
	// testing
	code, body := tester.New().RequestWithServe("GET", "/", nil, nil)

	// assertion
	assert.Equal(t, 200, code)
	assert.Equal(t, "welcome home", body)
}

func TestRequestWithHeadersTest(t *testing.T) {
	// testing
	headers := map[string]string{"KEY": "VALUE"}
	e, rec := tester.New().RequestWithContext("GET", "/", headers, nil)

	// assertion
	assert.NotNil(t, e)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}

func TestSetupHandlerTest(t *testing.T) {
	// testing
	h, ret := tester.New().SetupHandlerTest()

	// assertion
	assert.NotNil(t, h)
	assert.NotNil(t, ret)
}
