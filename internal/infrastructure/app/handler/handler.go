package handler

import (
	"github.com/shandysiswandi/echo-service/internal/config"
	"github.com/shandysiswandi/echo-service/internal/domain/port"
	"github.com/shandysiswandi/echo-service/pkg/clock"
	"github.com/shandysiswandi/echo-service/pkg/gen"
	"github.com/shandysiswandi/echo-service/pkg/validation"
)

/* Example get user jwt
user := c.Get("user").(*jwte.Token)
data := user.Claims.(*JWTClaim)
*/

type (
	Handler struct {
		config    *config.Config
		validate  *validation.Validation
		clock     clock.Clocker
		generator gen.Generator
		tdu       port.TodoUsecase
	}

	HandlerConfig struct {
		Config      *config.Config
		Validator   *validation.Validation
		Clock       clock.Clocker
		Generator   gen.Generator
		TodoUsecase port.TodoUsecase
	}
)

func New(hc HandlerConfig) *Handler {
	return &Handler{
		config:    hc.Config,
		validate:  hc.Validator,
		clock:     hc.Clock,
		generator: hc.Generator,
		tdu:       hc.TodoUsecase,
	}
}
