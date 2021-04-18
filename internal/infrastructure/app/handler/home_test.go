package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/util/response"
	"github.com/stretchr/testify/assert"
)

func Test_NewHomeHandler_Home(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.NewHomeHandler()

	// Assertions
	if assert.NoError(t, h.Home(c)) {
		res, err := response.SuccessForTest(rec.Body.String())
		assert.NoError(t, err)
		assert.NotNil(t, res)
		//
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, false, res.Error)
		assert.Equal(t, "welcome home", res.Message)
		assert.Equal(t, []interface{}{}, res.Data)
	}
}

func Test_NewHomeHandler_Graceful(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/graceful", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.NewHomeHandler()
	expBody := "\"OK\"\n"

	// Assertions
	if assert.NoError(t, h.Graceful(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expBody, rec.Body.String())
	}
}
