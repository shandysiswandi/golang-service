package tester

import (
	"io"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func RequestTest(e *echo.Echo, m, path string, body io.Reader) (int, string) {
	req := httptest.NewRequest(m, path, body)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

func RequestWithHeadersTest(e *echo.Echo, m, path string, headers map[string]string, body io.Reader) (int, string) {
	req := httptest.NewRequest(m, path, body)
	rec := httptest.NewRecorder()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}
