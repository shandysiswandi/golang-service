package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/util/response"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	cfg := &config.Config{JWTSecret: ""}
	e := app.New(cfg)

	c := e.NewContext(req, rec)

	hc := handler.HandlerConfig{}
	h := handler.New(hc)

	// Assertions
	assert.NoError(t, h.Home(c))
	res, err := response.SuccessForTest(rec.Body.String())
	assert.NoError(t, err)
	assert.NotNil(t, res)
	//
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, false, res.Error)
	assert.Equal(t, "welcome home", res.Message)
	assert.Equal(t, []interface{}{}, res.Data)
}
