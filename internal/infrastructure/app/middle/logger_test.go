package middle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/middle"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	// setup
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// testing
	h := middle.Logger()(func(c echo.Context) error {
		return c.String(http.StatusOK, "logger")
	})
	h(c)

	// assertion
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, "logger", rec.Body.String())
}
