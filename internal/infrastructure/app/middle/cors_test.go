package middle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/middle"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	// setup
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// testing
	req.Header.Set(echo.HeaderOrigin, "*")
	h := middle.CORS()(echo.NotFoundHandler)
	h(c)

	// assertion
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, "*", rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
}
