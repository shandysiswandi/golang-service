package tester

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/mongodb"
	"github.com/shandysiswandi/echo-service/mocks"
	"github.com/shandysiswandi/echo-service/pkg/validation"
)

type (
	ReturnSetupHandlerTest struct {
		Config      *config.Config
		Validation  *validation.Validation
		Generator   *mocks.Generator
		Clocker     *mocks.Clocker
		TodoUsecase *mocks.TodoUsecase
	}

	Tester interface {
		RequestWithServe(string, string, map[string]string, io.Reader) (int, string)
		RequestWithContext(string, string, map[string]string, io.Reader) (echo.Context, *httptest.ResponseRecorder)
		SetupHandlerTest() (*handler.Handler, *ReturnSetupHandlerTest)
	}

	tester struct{}
)

func New() Tester {
	return &tester{}
}

func (t *tester) reqres(m, path string, headers map[string]string, body io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(m, path, body)
	rec := httptest.NewRecorder()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req, rec
}

func (t *tester) setupApp() *echo.Echo {
	cfg := config.New()
	dbm := mongodb.New(cfg)
	e := app.New(cfg, dbm)
	return e
}

func (t *tester) RequestWithServe(m, path string, h map[string]string, body io.Reader) (int, string) {
	req, rec := t.reqres(m, path, h, body)
	e := t.setupApp()
	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

func (t *tester) RequestWithContext(m, path string, h map[string]string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	req, rec := t.reqres(m, path, h, body)
	e := t.setupApp()
	return e.NewContext(req, rec), rec
}

func (t *tester) SetupHandlerTest() (*handler.Handler, *ReturnSetupHandlerTest) {
	cfg := config.New()
	val := validation.New()
	gen := new(mocks.Generator)
	clk := new(mocks.Clocker)
	tdu := new(mocks.TodoUsecase)

	h := handler.New(handler.HandlerConfig{
		Config:      cfg,
		Validator:   val,
		Generator:   gen,
		Clock:       clk,
		TodoUsecase: tdu,
	})

	rsht := &ReturnSetupHandlerTest{
		Config:      cfg,
		Validation:  val,
		Generator:   gen,
		Clocker:     clk,
		TodoUsecase: tdu,
	}

	return h, rsht
}
